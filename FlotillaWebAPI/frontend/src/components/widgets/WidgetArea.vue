<template>
  <div class="widgetGrid" ref="grid">
    <template v-for="r in units.rows">
      <template v-for="c in units.cols">
        <div
          :key="r+'-'+c"
          :id="r+'-'+c"
          :style="{ top: ((r - 1) * units.v) + '%',
                left: ((c - 1) * units.h) + '%', 
                width: units.h + '%', 
                height: units.v + '%'}"
          :class="[(r+'-'+c) in gridlist ? 'hoverLoc' : '', 'gridLoc']"
          v-if="movingItem"
          @mouseover="CallMouseOver([r-1,c-1])"
        >({{r-1}}, {{c-1}})</div>
      </template>
    </template>
    <BaseWidget
      v-for="widget in activeWidgets"
      v-bind:key="widget.name"
      :Name="widget.name"
      :units="units"
      :position="position"
      :availablelocs="availableLocs"
      :unavailablelocs="unavalableLocs"
      :component="widget.widget"
      @drag="dragEvent"
      @resize="resizeEvent"
    />
  </div>
</template>

<script>
import BaseWidget from "@/components/widgets/BaseWidget";
import Vue from "vue";

export default {
  name: "WidgetArea",
  data() {
    return {
      movingItem: false,
      units: {
        v: 0,
        h: 0,
        rows: 10,
        cols: 10
      },
      hoverW: 0,
      hoverH: 0,
      position: [0, 0],
      gridlist: [],
      availableLocs: [],
      unavalableLocs: [],

    };
  },
  watch: {
    "units.v"() {
      this.CalculateGrid();
    },
    "units.h"() {
      this.CalculateGrid();
    },
    "this.$refs.grid.clientWidth"() {
      this.CalculateGrid();
    },
    "this.$refs.grid.clientHeight"() {
      this.CalculateGrid();
    },
    "this.$store.state.activeWidgets"(){
      this.UnavailableGridSpots();
      this.AvailableGridSpots();
    },
    "this.$store.state.coreWidgets"(){
      this.UnavailableGridSpots();
      this.AvailableGridSpots();
    }

  },
  methods: {
    CalculateGrid() {
      this.units.v = 100 / this.units.rows;
      this.units.h = 100 / this.units.cols;
    },

    dragEvent(event) {
      this.movingItem = true;
      this.hoverW = event.w + 1;
      this.hoverH = event.h + 1;
      document.documentElement.addEventListener(
        "mouseup",
        this.stopMoving,
        false
      );
    },
    resizeEvent(event) {
      this.movingItem = true;
      document.documentElement.addEventListener(
        "mouseup",
        this.stopMoving,
        false
      );
    },
    stopMoving() {
      this.movingItem = false;
      this.hoverW = 0;
      this.hoverH = 0;
      this.gridlist = [];
      document.documentElement.removeEventListener(
        "mouseup",
        this.stopMoving,
        false
      );
    },
    CallMouseOver(position) {
      this.position = position;

      this.gridlist = [];
      // Update Ids
      for (var i = position[0] + 1; i < position[0] + this.hoverH; i++) {
        for (var j = position[1] + 1; j < position[1] + this.hoverW; j++) {
          Vue.set(this.gridlist, i + "-" + j);
        }
      }
    },
    FindPosition(widgetName){
      var Coord = this.GetWidgetCoords(widgetName)

      if (this.IsCoordValid(Coord)){
        return([Coord.x, Coord.y])
      } else{
        return [0,0]
      }
    },
    IsCoordValid(Coord){
      var locs = this.MakeLocs(Coord)
      locs.forEach(loc => {
        if (!this.AvailableGridSpots.includes(loc)){
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
      for (var i = startingCoord[0]; i < startingCoord[0] + Coord.h; i++) {
        for (var j = startingCoord[1]; j < startingCoord[1] + Coord.w; j++) {
          used.push(i + "-" + j)
        }
      }
      return used
    },
    AvailableGridSpots: function(){
      var allSpots = []
      for (var i = 0; i < 10; i++) {
        for (var j = 0; j < 10; j++) {
          var loc = i + "-" + j
          if (!this.unavalableLocs.includes(loc)){
            allSpots.push(loc)
          }
        }
      }
      this.availableLocs = allSpots
    },
    UnavailableGridSpots: function(){
      var used = []
      this.activeWidgets.forEach(widget => {
        var Coord = this.GetWidgetCoords(widget.name)
        var componentlocs = this.MakeLocs(Coord)
        componentlocs.forEach(item =>{
          used.push(item)
        }) 
      })
      this.unavalableLocs = used
    }
  },
  created() {
    this.CalculateGrid();
    this.UnavailableGridSpots();
    this.AvailableGridSpots();
  },
  computed: {
    activeWidgets: function() {
      return this.$store.state.activeWidgets;
    },
    
  },
  components: {
    BaseWidget,
  }
};
</script>

<style scoped>
.gridLoc {
  position: absolute;
  color: azure;
  display: flex;
  flex-direction: column;
  justify-content: center;
  border: 1px solid grey;
  z-index: 3;
  transition: all 1s;
}

.gridLoc:hover {
  box-shadow: inset 0px 0px 4px 4px white;
  transition: all 0.4s;
}

.hoverLoc {
  box-shadow: inset 0px 0px 4px 4px white;
  transition: all 0.4s;
}

.widgetGrid {
  position: relative;
  height: 100%;
}
</style>

