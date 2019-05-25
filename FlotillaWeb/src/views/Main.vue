<template>
  <div class="layout">
    <TopBar class="fixedTop"/>
    <div class="remainingArea">
      <WidgetArea />
    </div>
    
    <!-- <transition-group name="fade">
    <div class="widgetArea" v-bind:key="widget" v-for="widget in activeWidgets">
      <div class="fader">
        <BaseWidget :Name="widget">
          {{widget}}
        </BaseWidget>
      </div>			
    </div>
  </transition-group> -->
  </div>
  
</template>

<script>
import TopBar from '@/components/TopBar'
//import BaseWidget from '@/components/widgets/BaseWidget'
import WidgetArea from '@/components/widgets/WidgetArea'

export default {
  name: 'FlotillaMain',
  components: {
    TopBar,
    //BaseWidget,
    WidgetArea,
  },
  computed: {
    activeWidgets: function(){
        return this.$store.state.activeWidgets
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.remainingArea{
  flex-grow: 1;
}

.fixedTop{
  height: 40px;
}

.layout{
  display: flex;
  flex-flow: column;
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
