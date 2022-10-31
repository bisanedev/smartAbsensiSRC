<template>
<modal name="semester-add" draggable=".card-header" :width="400" :height="400.9" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Semester</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('semester-add')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">  
      <div class="col-md-12">
        <label for="nama">Tahun Ajaran</label>
         <input type="text" class="form-control" placeholder="Tahun Ajaran" v-model="semester">
      </div>
      <div class="col-md-12">
        <label for="nama">Pilih Semester</label>
        <select class="form-control" v-model="semesterdetail">
          <option>Ganjil</option>
          <option>Genap</option>          
        </select>           
      </div>
      <div class="col-md-12">
        <label for="mapel">Tanggal Awal Semester</label>
        <date-picker v-model="waktustart" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal">                          
        </date-picker>     
      </div>
      <div class="col-md-12">
        <label for="mapel">Tanggal Akhir Semester</label>
        <date-picker v-model="waktuend" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal">                          
        </date-picker>       
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="disable" @click="addsemester()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">

export default {
  name: 'SemesterAddModal',
  data () {
    return { 
      disable: false,         
      semester:'',
      semesterdetail:'',
      waktustart:'',
      waktuend:'',
    }
  },
  methods: {      
    addsemester:function(event){
        if (this.semesterdetail == '' && this.semester == '' && this.waktustart == '' && this.waktuend == ''){
           this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: 'Periksa Kembali Input Data'
            });
          return
        }
        this.disable = true;
        axios.post(window.location.origin+'/api/semester', {
            nama: this.semester,semester:this.semesterdetail, waktustart: this.waktustart, waktuend: this.waktuend
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          {             
            //console.log('berhasil diupdate');
            this.$modal.hide('semester-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Semester Sukses',
              type: 'success',
              text: 'Data Semester '+this.semester+' Berhasil Di Tambahkan'
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
      this.semester = this.waktuend = this.waktustart = this.semesterdetail = '';     
    }
   //end method
  }
}
</script>