
<template>
    <v-container>
      <router-view></router-view> 
      <v-bottom-navigation fixed grow shift background-color="primary">
        <v-btn class="mx-3" fab color="secondary" v-on:click.native="RouterFiles">
          <v-icon>$vuetify.icons.solid_hdd</v-icon>
          <span> Files </span>
        </v-btn>
        <v-btn class="mx-3" fab color="secondary" v-on:click.native="RouterStats">
          <v-icon>$vuetify.icons.regular_chart_bar</v-icon>
          <span> Status </span>
        </v-btn>
        <v-btn class="mx-3" fab color="secondary" v-on:click.native="RouterUtil">
          <v-icon>$vuetify.icons.solid_hat_wizard</v-icon>
          <span> Util </span>
        </v-btn>
      </v-bottom-navigation>
      <!-- <FlotillaSideBar class="fixedTop"/>-->
    </v-container>
  
</template>
  

<script>
export default {
  name: 'FlotillaMain',
  computed: {
    activeWidgets: function(){
        return this.$store.state.activeWidgets
    },
  },
  methods: {
    RouterFiles (){
      console.log("pushed")
      this.$router.push("files")
    },
    RouterStats (){
      this.$router.push("status")
    },
    RouterUtil (){
      this.$router.push("util")

    }

  }
}
</script>

