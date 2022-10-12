<script setup>
  import axios from "axios";
  import { ref, unref, watch, computed} from "vue";
  import {useEmailStore} from '../../store/email.js';
  import router from "../../router";
  
  const emailStore = useEmailStore();

  //Variables no reactivas
  const timeout = 5000;
  const keys = ['From', 'To', 'Subject'];

  // Propiedades recibidas
  const props = defineProps({
    query: {
      type: [String],
      required: true
    }
  })
  
  //Variables reactivas
  const loadingEmails = ref(false);
  const errorEmails = ref(false)
  const items = ref([]);

  const totalItems = ref(1);
  const currentPage = ref(1);
  const itemsPerPage = ref(20);

  const bodySearch = ref({
    search_type: "matchall",
    query:  {
      term: ""
    }, 
    from: 0,
    max_results: 20,
    _source: []
  });

  bodySearch.value.max_results = itemsPerPage.value;

  //Propiedades computadas 
  const disabledScrollbar = computed(()=>{
    return (totalItems.value == items.value.length || errorEmails.value)
  })

  //Funciones
  const getEmails = async ()=>{
    return new Promise( async (resolve, reject)=>{
      try {
        if(loadingEmails.value) {
          reject(new Error("Emails are loading"));
          return;
        }

        loadingEmails.value = true;
        const resp = await axios.post("/zinc/correos/search/api", unref(bodySearch), { timeout })
        const hits = resp.data.hits;


        //Es la primera consulta de la query actual.
        if(items.value.length == 0)
          totalItems.value = hits.total.value;
        
        items.value.push(...hits.hits)

        //errorEmails depende de la respuesta asíncrona.
        if(hits)
          errorEmails.value = false;

      } catch (error) {      
        console.log(error);
        errorEmails.value = true;
        
      } finally {
        loadingEmails.value = false;
        if(errorEmails.value)
          reject(new Error("Error loading emails"))
        else {
          // add 1 to  currentPage.value.
          currentPage.value +=1;
          resolve(currentPage.value);
        } 
      }
    });
  }
  
  const onLoad = async (index, done)=> {
    bodySearch.value.from = (currentPage.value - 1) * itemsPerPage.value;
    try {
      //Add 1 to currenPage.value if request is success 
      await getEmails()
      done()
    
    } catch (error) {
      done()
    }
  }

  const encodeUTF8_To_B64 = (str) =>{
    return window.btoa(unescape(encodeURIComponent(str)));
  }

  const showEmail = (index)=>{
    const obj = items.value[index];
    emailStore.setEmail(obj);
    router.push(`/email/${encodeUTF8_To_B64(obj["_id"])}`)
  }

  //Observadores de variables reactivas
  watch(()=>props.query, (newQuery, oldQuery)=> {
    items.value = [];
    bodySearch.value.search_type = (newQuery != "") ?  "match" : "matchall";
    bodySearch.value.query.term = (newQuery != "" ) ? newQuery : "";

    currentPage.value = 1
    bodySearch.value.from = (currentPage.value - 1) * itemsPerPage.value;

    //Activa onLoad de forma automática (Obliga a cargar 1 página).
    if(totalItems.value==0)
      totalItems.value=1;
    else 
      onLoad(currentPage.value, ()=>{});  
  })
</script>

<template>
  <section class="section">
    <h2 class="section__title text-center">
      Listado de Correos 
      <q-badge class="q-p-xl q-mb-xl">
        {{totalItems}}
        <q-tooltip class="bg-indigo" :offset="[10, 10]">
          Total de correos encontrados
        </q-tooltip>
      </q-badge>
    </h2>
    <div class="q-pa-md listado">
      <q-infinite-scroll @load="onLoad" :offset="500" :debounce="300" :disable="disabledScrollbar">
        <div v-for="(item, index) in items" :key="index" v-memo="[index]" class="email" :class="{'email-even': index%2!=0, 'email-odd': index%2==0}">
          <p class="email__number">{{index + 1}}</p>
          <div class="email__content">
            <div class="email__content__key" v-for="(key, i) in keys" :key="`email-${index}-${key}-${i}`">
              <h3 class="key">{{key}}</h3>
              <p class="value">{{item["_source"][key]}}</p>
            </div>
          </div>
          <p class="email__show" @click.stop.prevent="showEmail(index)">
            Ver
          </p>
        </div>
        <template v-slot:loading>
          <div class="row justify-center q-my-md">
            <q-spinner-dots color="primary" size="40px" />
          </div>
        </template>
      </q-infinite-scroll>
      <div v-if="disabledScrollbar && errorEmails" class="q-my-lg text-center">
        <div v-if="loadingEmails" class="row justify-center">
          <q-spinner-dots color="primary" size="40px" />
        </div>
        <q-btn 
          v-else
          @click="errorEmails = false" 
          color="red" 
          icon-right="send"
          label="Error al cargar los correos, intente nuevamente" 
        />
      </div>
      <div v-if="disabledScrollbar && totalItems ==0">
        <h3 class="text-center text-warning">No se encontraron coincidencias en correos</h3>
      </div>
    </div>
  </section>
</template>

<style scope>
.section {
  width: auto;
  display: flex;
  flex-flow: column nowrap;
  align-items: stretch;
  justify-content: start;
  margin-top: 80px;
}

.section__title {
  font-size:3.5em;
  color: var(--color-icons)
}

h3 {
  font-size: 1.5rem;
}

.email {
  position: relative !important;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  width:100%;
  align-items: stretch;
  margin: 10px 0px !important;
  border-radius: 16px;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px, rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;;
}

.email-even {
  background-color: var(--color-secondary);
  color: white;
}

.email-odd {
  background-color: white;
  color:black
}

.email__number, .email__show {
  width: 8%;
  padding: 2%;
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-icons);
  color:white;
}

.email__number {
  border-top-left-radius: 16px;
  border-bottom-left-radius: 16px;
}

.email__show  {
  border-top-right-radius: 16px;
  border-bottom-right-radius: 16px;
  cursor: pointer;
}

.email__content {
  width: 80%;
  display: flex;
  flex-flow: column;
  align-items: stretch;
}

.email__content__key {
  display: flex;
  flex-wrap: wrap;
  align-content: stretch;
  justify-content: flex-start;
}

.email__content__key .key {
  width: 20%;
  font-size: 1.2rem;
  margin: 0;
} 

.email__content__key .value {
  width: calc(80% - 10px);
  margin-left: 10px;
  margin-bottom: 0;
  align-self: center;
}

@media (max-width: 600px) {
  .section {
    width: 98vw;
  }

  .section__title  {
    font-size: 1.5rem;
  }

  

  .email {
    justify-content: space-between;
    flex-flow: column wrap;
    align-items: stretch;
  }

  .email__number {
    width: unset;
    border-radius: 0px;
    border-top-left-radius: 16px;
    border-top-right-radius: 16px;
  }

  .email__show  {
    width: unset;
    border-radius: 0px;
    border-bottom-left-radius: 16px;
    border-bottom-right-radius: 16px;
  }

  .email__content__key {
    justify-content: stretch;
    align-items: flex-start;
  }

  .email__content__key .key, .email__content__key .value {
    width: unset;
  }

  h3 {
    font-size: 1.3rem !important;
  }
}
</style>