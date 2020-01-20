<template>
  <div class="about">
    <h1>Status</h1>
    <p>Status go here</p>
    <p>{{ Status }}</p>
    <input type='button' v-on:click="GetStatus" value="Get Status">
    <div>
      <textarea v-bind:value="CommOut" disabled rows="50" cols="150" ></textarea>
    </div>

  </div>
</template>

<script>
export default {
  name: 'FlotillaStatus',
  data(){
    return {
      Status: "No Status!",
      CommOut: "hello!\n"
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
      this.CommOut += line + "\n"
    }
  },
  created(){
    // Attach function to update when socket message comes in.
    flot_register_comm_callback(this.NewLineToComm)
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
</style>