<template>
  <div id="FlotillaSideBar" v-bind:style="SideBarStyle">
    <!-- Side Bar -->
    <div class="sidebar-container" >
      <ul>
        <li>
          <Hamburger :ButtonStyle="HamburgerStyle" 
                     v-bind:Action="MenuAction"
                     :collapseButton="showMenu"
                     :buttonText="ButtonTexts.hamburger">
            
              <Bars class="iconsize"/>
           
          </Hamburger>
        </li>
        <li>
          <Hamburger :ButtonStyle="HamburgerStyle" 
                     v-bind:Action="EmptyAction"
                     :collapseButton="showMenu"
                     :buttonText="ButtonTexts.file">
            <FilesIcon class="iconsize"/>
          </Hamburger>
        </li>
        <li>
          <Hamburger :ButtonStyle="HamburgerStyle" 
                     v-bind:Action="EmptyAction"
                     :collapseButton="showMenu"
                     :buttonText="ButtonTexts.status">
            <StatsIcon class="iconsize"/>
          </Hamburger>
        </li>
        <li>
          <Hamburger :ButtonStyle="HamburgerStyle" 
                     v-bind:Action="EmptyAction"
                     :collapseButton="showMenu"
                     :buttonText="ButtonTexts.util">
            <UtilIcon class="iconsize"/>
          </Hamburger>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import Bars from "@/assets/svg/solid/bars.svg"
  import StatsIcon from "@/assets/svg/solid/chart-line.svg"
  import FilesIcon from "@/assets/svg/solid/hdd.svg"
  import UtilIcon from "@/assets/svg/solid/hat-wizard.svg"
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
    Slide,
    Bars,
    StatsIcon,
    FilesIcon,
    UtilIcon
  },
  data () {
   return {
     showMenu: false,
     collapsed_width: 75,
     expanded_width: 200,
     SideBarStyle: {
       height: '100%',
       width: '75px'
     },
     HamburgerStyle:{
       backgroundColor: 'transparent', 
       color:'white',
       width: '100%',
       height: '100%'
     },
     ButtonTexts: {
       hamburger: 'Menu',
       file: "Files",
       status: "Status",
       util: "Util"

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
    EmptyAction (){
      console.log("Clicked!")
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#FlotillaSideBar{
  height: 100%;
  /* width: 75px; */
}

.sidebar-container{
  width: inherit;
  height: 100%;
  /* outline: 2px dashed blue; */
  border-right: 2px solid white;
}

ul{
  width: 100%;
  margin:0;
  padding:0;
}

li{
  min-height: 50px;
  max-height: 50px;
  width: 100%;
  margin: 0 0 25px 0;
}

.buttonrow{
  
  display: inline-block;

  
}

.iconsize{
  width: 45px;
  height: 45px;
  fill:white;
}


</style>
