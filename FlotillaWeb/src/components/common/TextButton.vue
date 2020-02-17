<template>
  <div style="width: inherit; height: inherit;">
    <input
      type="submit"
      v-bind:value="Label"
      v-on:click="Action"
      class="flotilla-button"
      v-bind:style="{backgroundColor: ButtonColor, color: TextColor, width:Width, height:Height}"
    >
  </div>
</template>

<script>
export default {
  name: "TextButton",
  props: {
    Label: {
      default: "No Value Assigned",
      type: String
    },
    ButtonColor: {
      default: "#73AD21",
      type: String
    },
    TextColor: {
      default: "white",
      type: String
    },
    Action: {
      default: function() {},
      type: Function
    },
    Width: {
      default: 100,
      type: Number
    },
    Height: {
      default: 100,
      type: Number
    }
  },
  components: {},
  methods: {}
};
</script>

<style scoped>
.flotilla-button {
  width: 100%;
  height: 100%;
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
</style>
