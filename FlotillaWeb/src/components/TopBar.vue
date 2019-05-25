<template>
  <div id="TopBar">
    <!-- Top Bar -->
    <div class="topbar-container">
      <div class="leftbar">
        <TextButton Label="Menu" ButtonColor="#2FA4F2" :Action="MenuAction"></TextButton>
      </div>
      <div class="centerLabel">Flotilla</div>
      <div class="rightbar">
        <TextButton ></TextButton>
      </div>
    </div>
    <LeftMenu :ShowMenu="showMenu" />
  </div>
</template>

<script>
  import TextButton from "./common/TextButton"
  import LeftMenu from "./LeftMenuComponents/LeftMenu"

  export default {
  name: 'TopBar',
  components: {
    TextButton,
    LeftMenu
  },
  data () {
   return {
     showMenu: false
   }
  },
  methods: {
    MenuAction () {
      this.showMenu = !this.showMenu
    },
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.centerLabel{
  display: flex;
  align-items: center;
  justify-content: center;
  width: calc(100% - (75px + 75px));


  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: White;
  font-size: 30px;

}

.buttonStyle{
  width: inherit;
  height: inherit;
}

.topbar-container{
  width: 100%;
  height: inherit;
  display: flex;
  flex-direction: row;
}
.rightbar{
  width: 75px;
  height: inherit;
  float: right;
}
.leftbar{
  width: 75px;
  height: inherit;
  float: left;
}

</style>
