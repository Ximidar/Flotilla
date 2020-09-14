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
    </v-tabs>
  </v-container>
</template>

<script>
import { Flotilla } from "@/flotilla"
import FlotConsole from "@/views/Status/console.vue"
import FlotControl from "@/views/Status/control/printer_control.vue"

export default {
  name: 'FlotillaStatus',
  components: {
    FlotConsole,
    FlotControl
  },
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
      this.Status = this.flot.GetStatus().then( (status) =>{
        if (!status){
          this.Status = "Idle"
          return
        }
        this.Status = status
      })
     
    },
    PostStatus () {
      this.flot.PostAction("Pause")
    },
    SendCancel () {
      this.flot.PostAction("Cancel")
    },
    NewLineToComm (line){
    //   this.CommOut.push(line)
    //   if (this.CommOut.length >= 200){
    //     var cut_point = this.CommOut.length - 200
    //     this.CommOut = this.CommOut.slice(cut_point)
    //   }
    //   var comm_wrapper = this.$el.querySelector("#comm-wrapper")
    //   comm_wrapper.scrollTop = comm_wrapper.scrollHeight
    },
    WS_Select (action){
      switch(action){
        case "NewStatus":
          this.GetStatus()
          break
        default:
          this.NewLineToComm(action)
          break
      }
    }
  },
  created(){
    var status_vue = this
    // Connect to WebSocket
    var ws = new WebSocket("ws://0.0.0.0:5000/api/ws")
    ws.onopen = function(){
      ws.send("Hello!")
    }
    ws.onmessage = function(evt){
      status_vue.WS_Select(evt.data)
    }

    // Setup flot
    this.flot = new Flotilla()
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