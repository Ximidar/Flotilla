<template>
  <div class="about">
    <h1>Status</h1>
    <p>Status go here</p>
    <p>{{ Status }}</p>
    <input type='button' v-on:click="GetStatus" value="Get Status">
    <div class="comm-wrapper" id="comm-wrapper">
      <span class="comm-item" v-for="message in CommOut">{{message}}</span>
    </div>

  </div>
</template>

<script>
export default {
  name: 'FlotillaStatus',
  data(){
    return {
      Status: "No Status!",
      CommOut: ["Hello!", "aloha", "comprender"]
    }
  },
  methods: {
    SetStatus(status){
      this.Status = status
    },
    GetStatus () {
      var status_vue = this
      function status (status) {
        status_vue.SetStatus("Got File info")
        status_vue.NewLineToComm(status)
        console.log(status)
      }
      flot_get_files(status)
    },
    NewLineToComm (line){
      this.CommOut.push(line)
      if (this.CommOut.length >= 200){
        var cut_point = this.CommOut.length - 200
        this.CommOut = this.CommOut.slice(cut_point)
      }
      var comm_wrapper = this.$el.querySelector("#comm-wrapper")
      comm_wrapper.scrollTop = comm_wrapper.scrollHeight
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
      status_vue.NewLineToComm(evt.data)
    }
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
</style>