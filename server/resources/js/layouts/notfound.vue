<template>
        <div class="container text-center" style="margin-top:160px;">
          <div class="display-1 text-muted mb-5"><i class="si si-exclamation"></i> 404</div>
          <h1 class="h2 mb-3">Oops.. You just found an error page..</h1>
          <p class="h4 text-muted font-weight-normal mb-7">We are sorry but our service is currently not available&hellip;</p>
          <button class="btn btn-primary" :disabled="disable" v-on:click="back()">
            <i class="fe fe-arrow-left mr-2"></i>Go back
          </button>
        </div>
</template>
<script type="application/javascript">
export default { 
 data () {
    return {
      disable:false,
    }
  }, 
  methods: {
    back: function () {      
      this.disable = true;
      this.$router.go(-1);   
    }
  }
// ...
}   
</script>