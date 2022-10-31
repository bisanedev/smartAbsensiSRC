<template>
	<div class="page">
      <div class="page-single">
        <div class="alert alert-warning" role="alert" style="width: 100%;position: absolute;top: 0;text-align: center;font-size: 12px;">                              
          Aplikasi promo/demo Absensi Sekolah 100% gratis dan tidak untuk di perjual belikan !! anda tertarik buat aplikasi bersama bisanedev hubungi kami di email: bisanedev@gmail.com , DM di instagram : instagram.com/bisanedev        
        </div>             
        <div class="container" style="margin-top:50px;">
          <div class="row">
            <div class="col col-login mx-auto" style="display: flex;flex-direction: column;align-items: center;">
              <div class="text-center mb-4">
                <img src="storage/logo.png" style="height: 160px;" alt="">
              </div>
              <form class="card" @submit.prevent="login" method="post" style="margin-bottom: 0;">
                <div class="card-body p-6">                  
                  <div class="form-group">
                    <label class="form-label">Username</label>
                    <input name="username" @keydown.space.prevent type="text" class="form-control" v-model="username" placeholder="Username">
                  </div>
                  <div class="form-group">
                    <label class="form-label">
                      Password                      
                    </label>
                    <input name="password" v-on:keyup.enter="login" type="password" class="form-control" v-model="password" placeholder="Password">
                  </div>                  
                  <div class="form-footer">
                    <button type="submit" v-on:keyup.enter="login" class="btn btn-primary btn-block">Sign in</button>
                  </div>
                </div>
              </form>
              <img src="assets/logo.png" style="width:180px;max-width: 100%;"/>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>
<script type="application/javascript">

export default {
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    login () {
      if(this.username === "" || this.password === ""){
        this.$notify({
            group: 'notifikasi',                  
            title: 'Failed Login',
            type: 'error',
            text: "Please fill Username Or Password !" 
           });
      }else{
      axios.post( window.location.origin + '/login' , {
        username: this.username , password: this.password
      })
      .then(response => {                
         if(response.data.status == "success")
         { 
          console.log('berhasil login');
          window.localStorage.setItem("token", response.data.token);
          window.localStorage.setItem("ipaddress", response.data.ipaddress);
          window.localStorage.setItem("warna", response.data.warna);  
          window.localStorage.setItem("sekolah", response.data.sekolah);  
          window.axios.defaults.headers.common['Authorization'] = response.data.token;
          this.$router.push({ name: 'dashboard' })
         }
        })
        .catch(error => {  
          if(error.response.status == 401){
           this.$notify({
            group: 'notifikasi',                  
            title: 'Failed Login',
            type: 'error',
            text: error.response.data.message 
           });
          }        
          console.log(error);
        })
      }
    }
    //-----
  }
  // ...
}
</script>