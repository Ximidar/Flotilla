<template>
  <div class="layout">
    <FlotillaSideBar class="fixedTop"/>
    <router-view style="{display: inline-block;}"></router-view>

  </div>
  
</template>

<script>
import FlotillaSideBar from '@/components/FlotillaSideBar'

export default {
  name: 'FlotillaMain',
  components: {
    FlotillaSideBar,
  },
  computed: {
    activeWidgets: function(){
        return this.$store.state.activeWidgets
    },
  },
  methods: {

  },
  created () {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.remainingArea{
  flex-grow: 1;
}

.fixedTop{
  /* height: 40px; */
  float: left;
  padding-right: 10px;
}

.layout{
  overflow: auto;
  height: inherit;
}

.fader{
  position: inherit;
  transition: transform 1s;
  width: fit-content;
  height: fit-content;
}

.fade-enter-active {
  opacity: 1;
  transition: all 0.5s;
}
.fade-leave-active {
  position: absolute;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
  transition: all 0.5s;
}
.fade-move{
  transition: transform 1s;
}
</style>
