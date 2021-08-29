<template>
  <div class="hamburger"
       
  >
    <button
      type="submit"
      v-on:click="Action"
      class="flotilla-button"
      v-bind:style="ButtonStyle"
    >
      <div class="butonalign">
        <div v-show="collapseButton" class="righttext">
          <div v-show="collapseButton" class="btext">{{ buttonText }}</div>
        </div>
        
        <div v-if="collapseButton" class="centerbutton">
          <!-- Icon -->
          <slot>
            <Bars class="iconsize"/>
          </slot>
        </div>
        <div v-else class="centerbuttoncollapsed">
          <!-- Icon -->
          <slot>
            <Bars class="iconsize"/>
          </slot>
        </div>

        </div>
      
    </button>
  </div>
</template>

<script> 
import Bars from "@/assets/svg/solid/bars.svg"
export default {
  name: "Hamburger",
  props: {
    ButtonStyle: {
      backgroundColor: "#73AD21",
      color: "white",
    },
    Action: {
      default: function(){},
      type: Function
    },
    collapseButton: {
      type: Boolean,
      default: false
    },
    buttonText: {
      default: "None!",
      type: String
    },
  },
  components: {
    Bars
  },
  data () {
    return {
    }
  },
  methods: {
  }
};
</script>

<style scoped>
.flotilla-button {
  border-radius: 6px;
  border: none;
  color: white;
  background-color: #73ad21;

  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
  
  
  
}

.flotilla-button:hover {
  opacity: 0.75;
}

.flotilla-button:focus {
  outline: none;
}

.flotilla-button:active {
  opacity: 0.3;
  width: 95%;
  height: 95%;
  -webkit-transition-duration: 0.2s;
  transition-duration: 0.2s;
}

.iconsize{
    width: 45px;
    height: 45px;
}

.buttonalign {
  text-align: 'center';
  overflow:auto;
}

.centerbutton{
  float:right;
  
}

.centerbuttoncollapsed {
  display:inline-block;
  
}

.righttext{
  position:relative;
  left: 0px;
  right:30px;
  top: 10px;
  display: inline-block;
  
}

.btext{
  width: 100%;
  height: 100%;
  color: whitesmoke;
  font-size: 25px;
  overflow: hidden;
  white-space: nowrap;
  
  
}

.hamburger{
  /* width: 100%; */
  height: 50px;
  max-height: 50px;
  width:100%;
  /* outline: 2px dashed rgb(231, 24, 24); */
}
</style>
