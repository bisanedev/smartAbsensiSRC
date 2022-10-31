<template>
<modal name="kelas-edit" draggable=".card-header" :width="400" :height="265.9" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Edit : <b>{{kelas}}</b></h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('kelas-edit')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">  
      <div class="col-md-6">
        <label for="urut">Nama Kelas</label>
         <input type="text" class="form-control" placeholder="Nama Kelas" v-model="kelas">
      </div>
      <div class="col-md-6">
        <label for="jumlah">Jumlah Siswa Aktif</label>
        <input type="text" class="form-control" placeholder="jumlah" v-model="jumlah" disabled>
      </div>
      <div class="col-md-6">
        <label for="mapel">Maksimum Mapel Harian</label>
        <div class="number-input">
          <button v-on:click="dec('maxpel', 1)" ></button>
          <input class="quantity" v-model.number="maxpel" type="number" style="width:70%;">
          <button v-on:click="inc('maxpel', 1)" class="plus"></button>
        </div> 
      </div>
      <div class="col-md-6">
        <label for="archive">Status</label>
        <select class="form-control" v-model="status" :disabled="disable">
          <option value="a">Aktif</option>
          <option value="n">Non-Aktif [arsipkan]</option>          
        </select>
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" @click="editkelas()">Update Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'KelasEditModal',
  data () {
    return {
      disable:true,
      id:'',
      kelas:'',
      status:'',
      jumlah:'',
      maxpel:0    
    }
  },
  watch: {
    jumlah() {
      if(this.jumlah === 0){        
         this.disable = false
      }
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
    fetchJumlahSiswa(){ 
    axios.get(window.location.origin+"/api/siswajumlah?kelas="+this.id)
      .then(response => {           
          this.jumlah = response.data.data          
        }).catch(error => this.kelasData = [])
    },  
    editkelas:function(event){
      
        axios.patch(window.location.origin+'/api/kelas/'+this.id, {
            kelas: this.kelas ,maxpel: this.maxpel, status: this.status,
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('kelas-edit')
            if(this.status === "a"){
                this.$notify({
                  group: 'notifikasi',                  
                  title: 'Update Kelas Sukses',
                  type: 'success',
                  text: 'Data Kelas '+this.kelas+' Berhasil Di Update'
                });
            }else{
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Kelas Sukses',
                type: 'warn',
                text: 'Data Kelas '+this.kelas+' Berhasil Di Arsipkan'
              });
            }            
          }
        })
        .catch(error => {        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })     
    },
    beforeOpen (event) {      
      this.id = event.params.wow.id;      
      this.kelas = event.params.wow.nama;
      this.status = event.params.wow.status;
      this.maxpel = event.params.wow.maxpel;
      this.fetchJumlahSiswa()
    },
    beforeClose (event) {
      //console.log("close")        
    }
   //end method
  }
}
</script>