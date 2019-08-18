<template>
  <div class="hamburger"
       
  >
    <button
      type="submit"
      v-on:click="Action"
      class="flotilla-button"
      v-bind:style="ButtonStyle"
    >
      <div class="centerbutton">
         <!-- Icon -->
        <Bars class="icon"/>
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
  },
  components: {
    Bars
  },
  methods: {}
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

.icon{
    width: 50px;
    height: 50px;
    float: right;
}

.centerbutton{
  height: 100%;
  width: 100%;
}

.hamburger{
  /* width: 100%; */
  height: inherit;
  width:100%;
  outline: 2px dashed rgb(231, 24, 24);
}
</style>
