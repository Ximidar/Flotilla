<template>
  <div class="Item">
    <div class="subItem">
        <IconBase width="30" height="30" icon-name="iconEnvelope" iconColor="white">
          <iconEnvelope />
        </IconBase>
    </div>
    <div class="subItem">{{ItemName}}</div>

    
      <div class="subItem" @click="addItem">
        <IconBase width="30" height="30" icon-name="iconEnvelope" iconColor="white" >
            <transition name="fade">
            <iconMinus v-if="ItemAdded"  />
            <iconPlus v-else  />
            </transition>
          </IconBase>
      </div>
    
  </div>
</template>



<script>
  import iconEnvelope from "@/components/icons/iconEnvelope"
  import iconPlus from "@/components/icons/iconPlus"
  import iconMinus from "@/components/icons/iconMinus"
  import IconBase from "@/components/icons/IconBase"

  export default{
    name: 'MenuItem',
    props: {
      ItemName: {
        default: 'Not Named',
        type: String
      },
      ItemSVGIcon: {
        default: 'none',
        type: String
      }
    },
    methods: {
      addItem() {
        this.$store.commit("toggleActiveWidget", this.ItemName)
      }
    },
    computed: {
      ItemAdded: function(){
        var name = this.ItemName
        var index = -1
        index = this.activeWidgets.findIndex(function(item){
          return item.name == name
        })
        if (index !== -1){
          return true
        }
        return false
      },
      activeWidgets: function(){
        return this.$store.state.activeWidgets
      }
    },
    components: {
      iconEnvelope,
      IconBase,
      iconPlus,
      iconMinus
    },

  }
</script>

<style scoped>

.Item{
  border-bottom: 2px solid green;

  width: 100%;
  height: 45px;

  display: flex;
    align-items: center;
    justify-content: center;
}

.Item:hover{
  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
  opacity: 0.8;
  box-shadow: 0 12px 16px 0 rgba(0,0,0,0.24), 0 17px 50px 0 rgba(0,0,0,0.19);
  
}

.subItem{
  width: calc(100% / 3);
  height: inherit;

  display: flex;
    align-items: center;
    justify-content: center;
    -webkit-user-select: none;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}

</style>