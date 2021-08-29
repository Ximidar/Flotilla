<template>
  <div class="ConsoleContainer">
    <textarea class="textbox" :value="gcodeBox" id="gcodeBox" readonly/>
    <div class="lineInput">
      <input class="gcode-input" type="text" v-on:keyup.enter="sendGcode" v-model="gcodeinput">
      <TextButton class="gcode-button" Label="Submit" :Action="sendGcode"/>
    </div>
  </div>
</template>

<script>
import TextButton from "@/components/common/TextButton";
export default {
  name: "pconsole",
  data() {
    return {
      gcodeBox: "",
      gBoxInput: {},
      gcodeinput: "",
    };
  },
  methods: {
    sendGcode() {
      this.gcodeBox += this.gcodeinput + '\n';
      this.gcodeinput = "";
      this.scrollGcodeToEnd()
    },
    scrollGcodeToEnd(){
      var scroll = this.$el.querySelector("#gcodeBox");
      scroll.scrollTop = scroll.scrollHeight        
    },
  },
  components: {
    TextButton
  }
};
</script>

<style scoped>
.ConsoleContainer {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

.textbox {
  height: 100%;
  flex-grow: 1;
  line-height: normal;

  vertical-align: top;
  resize: none;
  background-color: grey;
  color: white;
  border: 1px solid #888;
}

.lineInput {
  height: 25px;
  display: flex;
  flex-direction: row;
}

.gcode-input {
  flex-grow: 1;
}
.gcode-button {
  margin: auto;
}
</style>


