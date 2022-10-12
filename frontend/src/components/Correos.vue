<script setup>
  // import Filtros from "./Correos/Filtros.vue";
  import Listado from "./Correos/Listado.vue";
  
  import {ref} from "vue";
  const query = ref("");
</script>

<template>
  <main id="main" class="main">
    <div class="search">
      <q-input 
        bottom-slots 
        v-model="query"
        label-color="indigo-13" 
        label="Buscar en correos" 
        bg-color="white"
        rounded
        outlined="outlined"
        :debounce="1500"
      >
        <template v-slot:append>
          <q-icon v-if="query !== ''" name="close" @click="query = ''" size="sm" class="color-icons cursor-pointer"></q-icon>
          <q-icon name="search" style="" size="sm" class="color-icons"></q-icon>
        </template>
      </q-input>
    </div>
    <Listado :query="query"></Listado>
  </main>
</template>

<style scoped>
.main {
  max-width: 100vw;
  position: relative !important;
  display: flex;
  justify-content: center;
  background-color: rgba(203, 220, 247, 1.0);
}

.search {  
  position: absolute !important;
  width: 80vw;
  left:10vw;
  top: -25px;
  z-index: 1;
}
</style>
