<template>
  <v-container>
    <v-system-bar color="primary"
                  fixed
                  window
    >
      <v-icon>$vuetify.icons.solid_lock_open</v-icon>
      <v-icon>$vuetify.icons.solid_lock</v-icon>
      <span>{{ Status }}</span>
      <v-spacer></v-spacer>
      <v-icon>$vuetify.icons.solid_wifi</v-icon>
    </v-system-bar>
    <v-tabs>
      <v-tab
        href="#tab-flot-overview"
      >Overview</v-tab>
      <v-tab
        href="#tab-flot-console"
      >Console</v-tab>
      <v-tab
        href="#tab-flot-control"
      >Control</v-tab>
      <v-tab
        href="#tab-flot-temp"
      >Temperature</v-tab>
      <v-tab-item
        value="tab-flot-console"
      >
        <FlotConsole/>
      </v-tab-item>
      <v-tab-item
        value="tab-flot-temp"
      >
        <span>Aloha!</span>
      </v-tab-item>
      <v-tab-item
        value="tab-flot-control"
      >
        <FlotControl/>
      </v-tab-item>
      <v-tab-item
        value="tab-flot-overview"
      >
        <FlotOverview/>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>

<script>
import FlotConsole from "@/views/Status/console.vue"
import FlotControl from "@/views/Status/control/printer_control.vue"
import FlotOverview from "@/views/Status/overview/overview"
import flotilla from '@/flotilla'

export default {
  name: 'FlotillaStatus',
  components: {
    FlotConsole,
    FlotControl,
    FlotOverview
  },
  mixins: [flotilla],
  data(){
    return {
      Status: "No Status!",
      PauseButtonText: "Not Playing",
      Pause: "Pause",
      Resume: "Resume",
      Cancel: "Cancel",
      CommOut: ["Hello!", "aloha", "comprender"]
    }
  },
  methods: {
    GetStatus () {
      
      this.flotGetStatus().then( (status) =>{
        if (!status){
          this.Status = "Idle"
          return
        }
        this.Status = status
      })
     
    },
    PostStatus () {
      this.flotPostAction("Pause")
    },
    SendCancel () {
      this.flotPostAction("Cancel")
    }
  },
  created(){
    this.GetStatus()
  }
}
</script>

<style>
.about{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: White;
  
}

.comm-wrapper{
  display:block;
  border: 2px solid red;
  overflow-y: scroll;
  max-height: 600px
}

.comm-item{
  display: block;
  text-align: left;

}

.status{
  display: block;
  text-align: left;
  padding-left: 100px;
}

.status_buttons{
  display: block;
  text-align: left;
  padding-left: 100px;
  width: 200px;
  height: 50px;
}

.PauseButton{
}
</style>