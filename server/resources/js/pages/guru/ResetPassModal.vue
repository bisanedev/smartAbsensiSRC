<template>
<modal name="reset-pass" draggable=".card-header" :width="400" :height="265.8" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Change Password : <b>{{username}}</b></h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('reset-pass')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-12">
        <label for="password">Passward</label>
        <input type="password" class="form-control" placeholder="Password" v-model="password">
      </div>
      <div class="col-md-12">
        <label for="password">Re-passward</label>
        <input type="password" class="form-control" placeholder="Re-password" v-model="repassword">
      </div>   
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" @click="updatePass()">Change Password</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'ResetPassModal',
  data () {
    return {
      id:'',
      username:'',
      password:'',
      repassword:''
    }
  },
  methods: {      
    updatePass:function(event){
      
        axios.patch(window.location.origin+'/api/users/'+this.id+'/password', {
          password: this.password, repassword: this.repassword
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('reset-pass')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Password',
              type: 'success',
              text: 'Data User '+this.username+' Berhasil Di Update'
            });
          }
        })
        .catch(error => {
        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })    
    },
    beforeOpen (event) {   
      this.id = event.params.wow.id;            
      this.username = event.params.wow.username;
    },
    beforeClose (event) {
      //console.log("close")            
    }
   //end method
  }
}
</script>