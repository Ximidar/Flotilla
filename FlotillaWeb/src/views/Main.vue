
<template>
  <v-app>
    <router-view dark style="{display: inline-block;}"></router-view> 
    <v-bottom-navigation>
      <v-btn fab dark small color="primary" v-on:click.native="RouterFiles">
        <v-icon>$vuetify.icons.flot.files</v-icon>
        <span> Files </span>
      </v-btn>
      <v-btn fab dark small color="primary" v-on:click.native="RouterStats">
      </v-btn>
      <v-btn fab dark small color="primary" v-on:click.native="RouterUtil">
      </v-btn>
    </v-bottom-navigation>
    <!-- <FlotillaSideBar class="fixedTop"/>-->
    
  </v-app>
</template>
  

<script>
import FlotillaSideBar from '@/components/FlotillaSideBar'
export default {
  name: 'FlotillaMain',
  components: {
    FlotillaSideBar
  },
  computed: {
    activeWidgets: function(){
        return this.$store.state.activeWidgets
    },
  },
  methods: {
    RouterFiles (){
      console.log("pushed")
      this.$router.push("files")
    },
    RouterStats (){
      this.$router.push("status")
    },
    RouterUtil (){
      this.$router.push("util")

    }

  },
  created () {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped>

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

.iconsize{
  width: 45px;
  height: 45px;
  fill:white;
}
</style>
