<template>
<modal name="kelas-add" draggable=".card-header" :width="400" :height="265.9" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Kelas</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('kelas-add')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">  
      <div class="col-md-12">
        <label for="nama">Nama Kelas</label>
         <input type="text" class="form-control" placeholder="Nama Kelas" v-model="kelas">
      </div>
      <div class="col-md-12">
        <label for="mapel">Maksimum Mapel Harian</label>
        <div class="number-input">
          <button v-on:click="dec('maxpel', 1)" ></button>
          <input class="quantity" v-model.number="maxpel" type="number">
          <button v-on:click="inc('maxpel', 1)" class="plus"></button>
        </div>        
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="disable" @click="addkelas()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'kelasAddModal',
  data () {
    return { 
      disable: false,     
      kelas:'',
      maxpel:0
    }
  },
  methods: { 
    inc(property, amt){
      this[property] += amt
    },
    dec(property, amt){
      if (this[property] === 1) return
      this[property] -= amt
    },  
    addkelas:function(event){
        this.disable = true;
        axios.post(window.location.origin+'/api/kelas', {
            kelas: this.kelas, status: "a", maxpel: this.maxpel
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          {             
            //console.log('berhasil diupdate');
            this.$modal.hide('kelas-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Kelas Sukses',
              type: 'success',
              text: 'Data Kelas '+this.kelas+' Berhasil Di Tambahkan'
            });
          }
        })
        .catch(error => {   
            this.disable = false;     
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })     
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")
      this.maxpel = 0;
      this.kelas = '';      
    }
   //end method
  }
}
</script>