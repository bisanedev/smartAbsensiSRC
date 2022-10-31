<template>
	<modal name="user-upload" draggable=".card-header" :width="400" :height="432" @before-open="beforeOpen" @before-close="beforeClose">
	<div class="card">
	    <div class="card-header card-header-modal">
	     <h5 style="position: absolute; margin-top: 15px;">Foto : <b>{{nama}}</b></h5>
	     <div class="card-options">
	     	<button class="close" @click="$modal.hide('user-upload')">
	     		<i class="fe fe-x text-muted"></i>
	     	</button>
	 	 </div>
	    </div>
	    <div class="card-body card-body-modal">
	    <div class="form-row">
		  <croppa :height="resizableH" :initial-image="initialImage" :width="resizableW" class="resizable-croppa" v-model="croppa" style="display: block;margin: 0 auto;"></croppa>     
	    </div>
	      <div class="float-right" style="margin-top:20px;">
	        <button class="btn btn-primary" :disabled="disable" @click="updateFoto()">Upload Foto</button>
	      </div> 
	    </div>
	    </div>
	</modal>
</template>
<script type="application/javascript">
export default {
  name: 'UserUploadModal',
  data () {
    return {
      disable: false,
      resizableW: 200,
      resizableH: 300,
      id:'',
      username: '',
      nama:'',
      croppa: {},
      foto:''     
    }
  },
  computed:{
  	 initialImage() {
      return window.location.origin+'/storage/guru/'+this.foto
  	 }
  },
  methods: {
  	cleanString(string) {
  		if(string.length > 11) string = string.substring(0,11);
    	return string.replace(/[\s\/]/g, '')
	},         
    updateFoto(){
    if (!this.croppa.hasImage()) {        
       	this.$notify({
	        group: 'notifikasi',                  
	        title: 'Erorr No Image',
	        type: 'error',
	        text: 'No Image to Upload'
	    });                
        return
      }
    this.disable = true;
    this.croppa.generateBlob((blob) => {
    	const url = window.location.origin+'/api/upload/guru';
    	const config = { headers: { 'Content-Type': 'multipart/form-data' } };
        var fdata = new FormData()
        fdata.append('files', blob, this.cleanString(this.username)+"_"+this.id+".jpg")
        fdata.append('id', this.id)

	      axios.post(url , fdata , config)
	      .then(response => {
	        this.$parent.fetchData();
	        if(response.data.status == "success")
	        { 
	          //console.log('berhasil diupdate');
	          this.$notify({
	            group: 'notifikasi',                  
	            title: 'Update Success',
	            type: 'success',
	            text: 'Foto User '+this.username+' Berhasil Di Update'
	          });
	          this.$modal.hide('user-upload')
	        }
	      })
	      .catch(error => {
	      	 this.disable = false;
	         console.log(error);
	      })
      })
      
    },
    beforeOpen (event) {   
      this.id = event.params.wow.id;      
      this.nama = event.params.wow.nama;
      this.username = event.params.wow.username;
      this.foto = event.params.wow.foto;  
    },
    beforeClose (event) {
      //console.log("close")
      this.disable = false;
    }
  }
}
</script>