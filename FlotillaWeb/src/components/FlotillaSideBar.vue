<template>
  <div id="FlotillaSideBar" v-bind:style="SideBarStyle">
    <!-- Side Bar -->
    <div class="sidebar-container" >
      <div class="hamburger">
        <Hamburger ButtonColor="#2FA4F2" 
                   v-bind:Action="MenuAction">
        </Hamburger>
      </div>
    </div>
  </div>
</template>

<script>
  // import TextButton from "@/components/common/TextButton"
  import Hamburger from "@/components/common/Hamburger"
  import TextButton from "./common/TextButton"
  import { Slide } from 'vue-burger-menu'
  import Vue from 'vue'
  export default {
  name: 'FlotillaSideBar',
  components: {
    TextButton,
    Hamburger,
    Slide
  },
  data () {
   return {
     showMenu: false,
     collapsed_width: 75,
     expanded_width: 200,
     SideBarStyle: {
       height: '100%',
       width: '75px'
     }
   }
  },
  methods: {
    MenuAction () {
      this.showMenu = !this.showMenu
      if (this.showMenu){
        Vue.set(this.SideBarStyle, 'width', this.expanded_width + 'px')
      } else {
        Vue.set(this.SideBarStyle, 'width', this.collapsed_width + 'px')
      }
    },
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#FlotillaSideBar{
  height: 100%;
  /* width: 75px; */
}

.buttonStyle{
  width: inherit;
  height: inherit;
}

.sidebar-container{
  width: inherit;
  height: 100%;
  display: flex;
  flex-direction: row;
  outline: 2px dashed blue;
}


</style>
