<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Data Kelas</h1>
        </div>
        <div class="row row-cards row-deck">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <button type="button" class="btn btn-primary" @click="showAddModal()">
                            <i class="fe fe-plus"></i>
                        </button>
                        <div class="card-options">
                            <select class="form-control" v-model="perPage" @change="onChange($event)" style="width: 60px;margin-right:10px;">
                                <option>3</option>
                                <option>6</option>
                                <option>12</option>
                            </select>
                            <div class="input-group" style="width:260px;">
                            <input type="text" class="form-control" v-model="search" v-on:keyup="searchEnter" placeholder="Search Kelas">
                            <button class="btn bg-transparent" @click="refresh()" style="margin-left: -40px; z-index: 100;">
                                <i class="fe fe-x text-muted"></i>
                            </button>
                            <div class="input-group-append">
                                <button class="btn btn-primary" type="button" @click="cari()">
                                    <i class="fe fe-search"></i>
                                </button>
                            </div>
                            </div>
                        </div>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                        <div class="table-responsive">
                            <table class="table table-striped table-vcenter text-nowrap card-table">
                                <thead class="thead-dark">
                                    <tr>
                                        <th class="text-center w-1">#</th>     
                                        <th class="text-center" v-for="column in columns" :key="column" @click="sortByColumn(column)">
                                            {{ column | columnHead }} Kelas
                                        <span v-if="column === sortedColumn">
                                        <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                        <i v-else class="fe fe-arrow-down"></i>
                                        </span>
                                        </th>
                                        <th class="text-center" @click="sortByColumn('maxpel')">
                                        Maksimum Mata Pelajaran Harian
                                        <span v-if="'maxpel' === sortedColumn">
                                        <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                        <i v-else class="fe fe-arrow-down"></i>
                                        </span>
                                        </th>                                           
                                        <th class="text-center"></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr class="" v-if="tableData.length === 0">
                                        <td class="lead text-center" :colspan="columns.length + 3">No data found.</td>
                                    </tr>
                                    <tr v-else v-for="(data, key1) in tableData" :key="data.id">
                                        <td class="text-center">{{ serialNumber(key1) }}</td>
                                        <td class="text-center">         
                                        {{ data.nama }}
                                        </td>
                                        <td class="text-center">         
                                        {{ data.maxpel }}
                                        </td>
                                        <td class="text-center">
                                            <div class="item-action dropdown">
                                                <a href="javascript:void(0)" data-toggle="dropdown" class="icon"><i class="fe fe-grid"></i></a>
                                                <div class="dropdown-menu dropdown-menu-right">    
                                                    <button type="button" @click="showEditModal(data)" class="dropdown-item"><i class="fe fe-edit"></i> Edit 
                                                    </button>
                                                    <button class="dropdown-item" @click="showDeleteDialog(data.id,data.nama)"><i class="fe fe-trash"></i> Delete </a> </button>                                       
                                                </div>
                                            </div>
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                    <div class="card-footer">
                      <span v-if="pagination.total_record" style="position: absolute;bottom: 20px;"> &nbsp; <i>Total Data : {{ pagination.total_record }} entries.</i></span>
                        <nav v-if="pagination && tableData.length > 0">
                        <ul class="pagination float-right">
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(1)">
                                <i class="fe fe-chevrons-left"></i>
                              </a>
                            </li>
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(currentPage - 1)"><i class="fe fe-chevron-left"></i></a>
                            </li>                            
                            <li v-for="page in pagesNumber" v-if="Math.abs(page - currentPage) < 3 || page == pagesNumber - 1 || page == 0" class="page-item" :class="{'active': page == pagination.page}">
                            <a href="javascript:void(0)" @click.prevent="changePage(page)" class="page-link">{{ page }}</a>
                            </li>                            
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(currentPage + 1)"><i class="fe fe-chevron-right"></i></a>
                            </li>
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(pagination.total_page)"><i class="fe fe-chevrons-right"></i></a>
                            </li>
                        </ul>
                        </nav>
                    </div>
                </div>
            </div>
        </div>        
        <kelas-add-modal/>
        <kelas-edit-modal/>        
    </div>
</template>
<script type="application/javascript">

import KelasAddModal from './KelasAddModal.vue'
import KelasEditModal from './KelasEditModal.vue'

export default {
  components : {
    KelasAddModal,
    KelasEditModal
  },
  data() {
    return {      
      columns: ['nama'],
      search:'',
      tableData: [],
      url: window.location.origin,
      pagination:[],
      offset: 4,
      currentPage: 1,
      perPage: '3',
      sortedColumn: 'nama',
      order: 'desc'
    }
  },
  created() {
    return this.fetchData()
  },
  computed: {
    /**
     * Get the pages number array for displaying in the pagination.
     * */
    pagesNumber() {
      if (this.pagination.offset == undefined || this.pagination.offset == null) {
        return []
      }
      let from = this.pagination.page - this.offset
      if (from < 1) {
        from = 1
      }
      let to = from + (this.offset * 2)
      if (to >= this.pagination.total_page) {
        to = this.pagination.total_page
      }
      let pagesArray = []
      for (let page = from; page <= to; page++) {
        pagesArray.push(page)
      }
      return pagesArray
    },
    /**
     * Get the total data displayed in the current page.
     * */
    totalData() {
      return (this.pagination.offset - this.pagination.total_record) + 1
    }
  },
  methods: {
    fetchData() {
      let dataFetchUrl = `${this.url}/api/kelas?page=${this.currentPage}&search=${this.search}&order=${this.sortedColumn}&by=${this.order}&limit=${this.perPage}`
        axios.get(dataFetchUrl)
        .then(response => { 
          this.pagination = response.data
          this.tableData = response.data.records
        }).catch(error => this.tableData = [])
    },
    /**
     * Get the serial number.
     * @param key
     * */
    serialNumber(key) {
      return (this.currentPage - 1) * this.perPage + 1 + key
    },
    /**
     * Change the page.
     * @param pageNumber
     */
    changePage(pageNumber) {
      this.currentPage = pageNumber
      this.fetchData()
    },    
    /** perpage ***/
    onChange(event) {
      this.currentPage = 1;
      this.fetchData();
    },
    /** enter search **/
    searchEnter:function(e){
      if (e.keyCode === 13) {
        this.currentPage = 1;
        this.fetchData() 
      }
    },
    /** cari **/
    cari(){
      this.currentPage = 1;
      this.fetchData()
    },
    /** refresh **/
    refresh(){
      this.search = "";
      this.fetchData();     
    },    
    /** User ADD **/
    showAddModal(){
      this.$modal.show('kelas-add');
    },
    /** edit modal **/
    showEditModal(data){      
      this.$modal.show('kelas-edit',{wow:data});
    },
    /** pass variable to modal **/ 
    showDeleteDialog (id,nama) {
      this.$modal.show('dialog', {
        title: 'Konfirmasi',
        text: 'Do you really want to delete? Kelas : '+nama,
        buttons: [
          {
            title: 'Cancel',            
            handler: () => {
              this.$modal.hide('dialog')
            }
          },
          {
            title: 'Delete',
            default: true,
            handler: () => {                          
              axios.delete(this.url+'/api/kelas/'+id)
                .then(response => {                
                  if(response.data.status == "success")
                  { 
                    //console.log('berhasil dihapus');
                       
                    this.fetchData();
                    this.$modal.hide('dialog');
                    this.$notify({
                      group: 'notifikasi',                  
                      title: 'Delete Success',
                      type: 'success',
                      text: 'Data Kelas '+nama+' Berhasil Di Hapus'
                    });  
                  }
                })
                .catch(error => {
                    console.log(error);
                })
              //-------
            }
          }
        ]
      })
    },
    /**
     * Sort the data by column.
     * */
    sortByColumn(column) {
      if (column === this.sortedColumn) {
        this.order = (this.order === 'asc') ? 'desc' : 'asc'
      } else {
        this.sortedColumn = column
        this.order = 'asc'
      }
      this.fetchData()
    }
  },
  filters: {
    columnHead(value) {
      return value.split('_').join(' ').toUpperCase()
    }
  } 


//end script  
}
</script>