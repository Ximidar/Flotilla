<template>
  <div class="widget" v-bind:style="style">
    <div class="widgetTop" id="widgetTop">
      <div class="fill" @mousedown.left.prevent.stop="DragEvent('drag', $event)">{{Name}}</div>
    </div>
    <div class="lineBorder"></div>
    <div class="slottedWidget">
      <component :is="component" />
    </div>
    <div class="resize-handle" @mousedown.left.prevent.stop="ResizeEvent('resize', $event)"></div>
  </div>
</template>

<script>
export default {
  name: "BaseWidget",
  props: ["Name", "units", "position", "component", "availablelocs", "unavailablelocs"],
  data() {
    return {
      ShowSettings: false,
      Moving: false,
      movingX: 0,
      movingY: 0,
      Resizing: false,
      resizingW: 0,
      resizingH: 0
    };
  },
  components: {},
  methods: {
    openSettings() {
      this.ShowSettings = !this.ShowSettings;
    },
    DragEvent(name, e) {
      this.$emit(name, {
        Name: this.Name,
        w: this.SizeW,
        h: this.SizeH,
        e: e
      });
      this.Moving = true;
      this.PosX = e.clientX;
      this.PosY = e.clientY;
      document.documentElement.addEventListener(
        "mousemove",
        this.MouseMover,
        false
      );
      document.documentElement.addEventListener("mouseup", this.MouseUp, false);
    },
    MouseMover() {
      this.PosX = this.position[1];
      this.PosY = this.position[0];
    },
    MouseUp() {
      document.documentElement.removeEventListener(
        "mousemove",
        this.MouseMover,
        false
      );
      document.documentElement.removeEventListener(
        "mouseup",
        this.MouseUp,
        false
      );
      // Update X, Y
      var updateCoord = {
        Name: this.Name,
        Coord: this.position[1],
        elementID: "x"
      };
      this.$store.commit("updateWidgetCoord", updateCoord);
      updateCoord = {
        Name: this.Name,
        Coord: this.position[0],
        elementID: "y"
      };
      this.$store.commit("updateWidgetCoord", updateCoord);

      this.Moving = false;
    },
    // Resize Helpers
    ResizeEvent(name, e) {
      this.$emit(name, {
        Name: this.Name,
        w: this.SizeW,
        h: this.SizeH,
        e: e
      });
      this.Resizing = true;
      this.position[0] = this.Sizeh;
      this.position[1] = this.SizeW;
      document.documentElement.addEventListener(
        "mousemove",
        this.ResizeMover,
        false
      );
      document.documentElement.addEventListener(
        "mouseup",
        this.ResizeUp,
        false
      );
    },
    ResizeMover() {
      this.SizeW =
        this.position[1] +
        1 -
        this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["x"];
      this.SizeH =
        this.position[0] +
        1 -
        this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["y"];
    },
    ResizeUp() {
      document.documentElement.removeEventListener(
        "mousemove",
        this.ResizeMover,
        false
      );
      document.documentElement.removeEventListener(
        "mouseup",
        this.ResizeUp,
        false
      );
      // Update W, H
      var updateCoord = {
        Name: this.Name,
        Coord: this.SizeW,
        elementID: "w"
      };
      this.$store.commit("updateWidgetCoord", updateCoord);
      updateCoord = {
        Name: this.Name,
        Coord: this.SizeH,
        elementID: "h"
      };
      this.$store.commit("updateWidgetCoord", updateCoord);

      this.Resizing = false;
    },
    // ########################################
    // # Calculate Available Spawning Positions
    // ########################################
    FindPosition(widgetName){
      var Coord = this.GetWidgetCoords(widgetName)

      if (this.IsCoordValid(Coord)){
        return([Coord.x, Coord.y])
      } else{
        return this.FindFirstValidPos
      }
    },
    FindFirstValidPos(){
      for (var i = 0 ; i < 10; i++) {
        for (var j = 0 ; j < 10; j++) {
          var Coord = new this.$store.Coord(i, j, this.SizeW, this.SizeH)
          if (this.IsCoordValid(Coord)){
            return [Coord.x, Coord.y]
          }
        }
      }

      return [5,5]
    },
    IsCoordValid(Coord){
      var locs = this.MakeLocs(Coord)
      locs.forEach(loc => {
        if (!this.availablelocs.includes(loc)){
          return false
        }
      });

      return true
      
    },
    GetWidgetCoords(widgetName){
      var index = -1
      index = this.$store.state.coreWidgets.findIndex(function(item){
        return item.ItemName == widgetName
      })
      if (index !== -1){
        return this.$store.state.coreWidgets[index].Coord
      }
      // We should never get here, but just incase
      return new this.$store.Coord(1,1,1,1)
    },
    MakeLocs(Coord){
      var used = []
      var startingCoord = [Coord.x, Coord.y]
      // Update locs
      for (var i = startingCoord[0] + 1; i < startingCoord[0] + Coord.h; i++) {
        for (var j = startingCoord[1] + 1; j < startingCoord[1] + Coord.w; j++) {
          used.push(i + "-" + j)
        }
      }
      return used
    }
  },
  created(){
    var startingPos = this.FindPosition(this.Name)

    // Update X, Y
    var updateCoord = {
      Name: this.Name,
      Coord: startingPos[0],
      elementID: "x"
    };
    this.$store.commit("updateWidgetCoord", updateCoord);
    updateCoord = {
      Name: this.Name,
      Coord: startingPos[1],
      elementID: "y"
    };
    this.$store.commit("updateWidgetCoord", updateCoord);

  },
  computed: {
    SizeW: {
      get: function() {
        if (this.Resizing) {
          return this.resizingW;
        }
        return this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["w"];
      },
      set: function(newW) {
        if (newW <= 0) {
          return this.resizingW;
        }
        this.resizingW = newW;
        return newW;
      }
    },
    SizeH: {
      get: function() {
        if (this.Resizing) {
          return this.resizingH;
        }
        return this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["h"];
      },
      set: function(newH) {
        if (newH <= 0) {
          return this.resizingH;
        }
        this.resizingH = newH;
        return newH;
      }
    },
    PosX: {
      get: function() {
        if (this.Moving) {
          return this.movingX;
        }
        return this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["x"];
      },
      set: function(newx) {
        this.movingX = newx;
        return newx;
      }
    },
    PosY: {
      get: function() {
        if (this.Moving) {
          return this.movingY;
        }
        return this.$store.state.coreWidgets.filter(
          widget => widget.ItemName == this.Name
        )[0].Coord["y"];
      },
      set: function(newy) {
        this.movingY = newy;
        return newy;
      }
    },
    style() {
      if (this.Moving) {
        return {
          top: this.PosY * this.units.v + "%",
          left: this.PosX * this.units.h + "%",
          width: this.SizeW * this.units.v + "%",
          height: this.SizeH * this.units.h + "%"
        };
      } else {
        return {
          top: this.PosY * this.units.v + "%",
          left: this.PosX * this.units.h + "%",
          width: this.SizeW * this.units.v + "%",
          height: this.SizeH * this.units.h + "%"
        };
      }
    }
  }
};
</script>

<style scoped>
.widget {
  border: 4px solid #2fa4f2;
  border-top-right-radius: 25px;
  border-top-left-radius: 25px;
  border-bottom-right-radius: 6px;
  border-bottom-left-radius: 6px;
  display: flex;
  flex-flow: column;
  transition: all 0.1s;
  position: absolute;
  z-index: 2;
}

.widgetTop {
  display: flex;
  flex-direction: row;
  width: calc(100% - 25px);
  margin: auto;
  height: 25px;
  padding-top: 5px;
}

.fill {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: move;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: white;
}

.slottedWidget {
  flex-grow: 1;
  align-items: center;
  justify-content: center;
  padding: 2px;
  color: white;
}

.lineBorder {
  border-bottom: 2px solid #2fa4f2;
  width: 100%;
  height: 2px;
}

/* resize */
.resize-handle {
  position: absolute;
  right: 0px;
  bottom: 0px;
  width: 5px;
  height: 5px;
  cursor: nwse-resize;
  border: 1px solid #2fa4f2;
  border-radius: 1.5px;
  background: #2fa4f2;
}

.resize-handle:hover {
  background-color: rgba(128, 128, 128, 0.1);
}
</style>