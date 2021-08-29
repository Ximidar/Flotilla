<template>
  <transition name="fade">
    <div v-if="ShowMenu" class="menu">
      <div v-bind:key="menuItem.ItemName" v-for="menuItem in MenuItems">
        <MenuItem  :ItemName="menuItem.ItemName" 
                   :ItemSVGIcon="menuItem.ItemSVGIcon"
          />
        </div>

    </div>
  </transition>
</template>
  
<script>
  import MenuItem from "./MenuItem"
export default {
  name: "LeftMenu",
  props: {
    ShowMenu: {
      default: false,
      type: Boolean
    },

  },
  computed: {
            MenuItems () {
                return this.$store.getters.getCoreWidgets
            }
        },
  components: {
    MenuItem,
  }
}
</script>

<style scoped>
.menu{
  margin-top: 10px;
    height: auto;
  width: 300px;

  overflow: hidden;

  border-radius: 25px;
  border: 5px;
  border-color: green;
  border-style: solid;

  color: white;
  z-index: 100;
  position: absolute;
  background-color: #031E28;

  display: flex;
  flex-direction: column;
  
}

.fade-enter-active, .fade-leave-active {
  opacity: 1;
  transition: all 0.5s
}
.fade-enter, .fade-leave-to {
  opacity: 0;
  width: 0;
  transition: all 0.5s
}
</style>