<template>
  <div class="about">
    <h1>Status</h1>
    <p>Status go here</p>
    <p>{{ Status }}</p>
    <input type='button' v-on:click="GetStatus" value="Get Status">
  </div>
</template>

<script>
export default {
  name: 'FlotillaStatus',
  data(){
    return {
      Status: "No Status!"
    }
  },
  methods: {
    SetStatus(status){
      this.Status = status
    },
    GetStatus () {
      var status_vue = this
      function status (status) {
        status_vue.SetStatus(status)
        console.log(status)
      }
      flot_get_files(status)
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
</style>