<script setup>
  import router from "../router/index.js";
  import { useEmailStore } from "../store/email";
  import { ref, nextTick } from "vue" 
  import axios from "axios";
  
  //Store
  const emailStore = useEmailStore();

  //Variables constantes
  const timeout = 10000;
  const fieldsEmail = [
    "Date",
    "From",
    "To",
    "Subject",
    "X-From",
    "X-To",
    "X-Origin",
    "X-bcc",
    "X-cc",
    "Content"
  ]

  //Variables reactivas
  const email = ref(null);
  const errorEmail = ref(false);
  const loadingEmail = ref(false);
  const id = ref(router.currentRoute.value.params.id);
  id.value = decodeURIComponent(atob(id.value));

  //Funciones
  const goBack = ()=>{
    router.push("/");
  }
  
  const getData = async ()=>{
    //Loading item from store.
    if(emailStore.getEmail)  {
      nextTick(()=>{
        email.value = emailStore.getEmail;
      });
      return;
    }
    
    //Loading item from backend
    try {
      //Email is getting
      if(loadingEmail.value) return;

      loadingEmail.value = true;

      const obj = {
        "from": 0,
        "max_results": 1,
        "query": {
          "bool": {
            "filter": [ {"term" : { "_id" : id.value} } ]
          }
        },
        "_source": []
      };

      const resp = await axios.post("/zinc/correos/search/es", obj, { timeout })
      const hits = resp.data.hits;
      email.value = hits.hits[0];
      //errorEmail depende de la respuesta asíncrona.
      if(resp)
        errorEmail.value = false;

    } catch (error) {
      console.log(error);
      errorEmail.value = true;
    } finally {
      setTimeout(()=>{
        loadingEmail.value = false;
      }, 100)
    }
  }

  getData();
</script>
<template>
  <main class="main bg-grey-5">
    <div class="main__btnGoBack">
      <q-btn
      rounded
      color="primary"
      size="md"
      label="Go Back"
      @click.stop.prevent="goBack"
      />
    </div>
    <q-card class="card_email">
      <q-card-section>
        <div v-if="email && email['_source']">
          <template v-for="( key, index) in fieldsEmail" :key="`${key}-${index}`">
            <div v-if="email['_source'][key]">
              <h3>{{key}}</h3>
              <p class="q-mb-lg">{{ email['_source'][key] }}</p>
            </div>
          </template>
        </div>
        <div v-else-if="loadingEmail" class="row justify-center q-my-md">
          <q-spinner-dots color="primary" size="40px" />
        </div>
        <div v-else-if="errorEmail" class="text-center">
          <q-btn 
            @click="getData" 
            color="red" 
            icon-right="send"
            label="Error al cargar el correo, intente nuevamente" 
          />
        </div>
        <div v-else>
          <h3 class="text-center text-warning">El correo que intenta ver no existe.</h3>
        </div>
      </q-card-section>
    </q-card>
  </main>
</template>

<style scoped>
.main {
  position: relative;
  display: flex;
  flex-flow: column wrap;
  justify-content: center;
  align-items: center;
  max-width:100vw;
  height:80vh;
}

.main__btnGoBack {
  position: absolute;
  left: 5px;
  top: -18px;
}

.card_email {
  max-height: 90%;
  max-width:85%;
  overflow-y: auto;
}

h3 {
  font-size: 1.5rem;
}

@media (max-width: 600px) {
  h3 {
    font-size: 1.3rem !important;
  }
}
</style>
