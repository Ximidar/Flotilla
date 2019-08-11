<template>
  <div style="width: 100%; height: 100%;">
    <button
      type="submit"
      v-on:click="Action"
      class="flotilla-button"
      v-bind:style="{backgroundColor: ButtonColor, color: TextColor}"
    >
        <!-- Icon -->
        <img class="img-icon" src="@/assets/svg/solid/bars.svg">
    </button>
  </div>
</template>

<script> 
export default {
  name: "Hamburger",
  props: {
    ButtonColor: {
      default: "#73AD21",
      type: String
    },
    TextColor: {
      default: "white",
      type: String
    },
    Action: {
      default: function(){},
      type: Function
    },
  },
  components: {

  },
  methods: {}
};
</script>

<style scoped>
.flotilla-button {
  width: inherit;
  height: inherit;
  border-radius: 6px;
  border: none;
  color: white;
  background-color: #73ad21;

  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
}

.flotilla-button:hover {
  opacity: 0.75;
  box-shadow: 0 12px 16px 0 rgba(0, 0, 0, 0.24),
    0 17px 50px 0 rgba(0, 0, 0, 0.19);
  /*transform: translateY(2px);*/
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

.img-icon{
    width:100%;
    height: 100%;
}
</style>
