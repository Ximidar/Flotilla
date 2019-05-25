<template>
  <div id="app">
    <router-view/>
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div> -->
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #031E28;
  padding-left: 5px;
  padding-right: 10px;
  padding-bottom: 5px;
  padding-top: 1px;
  height: 100%;
}

#nav {
  float: left;
  position: fixed;
  bottom: 10px;
  left: 10px;

}
html, body {
  height: 100%;
  max-width: 1500px;
  margin: auto;
  padding: 0;
  background: #031E28;
}
</style>
