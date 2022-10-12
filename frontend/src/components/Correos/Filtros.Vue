<!-- Idea que se quedó al implementar filtros en la búsqueda de correos -->
<script setup>
  import { ref } from "vue"
  const f = ref([])
  
</script>

<template>
  <Teleport :disabled="true">
    <section class="">
    <h2 class="section__title">Buscar en Campos</h2>
      <div class="container__filtros">
      </div>
    </section>
  </Teleport>
</template>
<style scoped>
@media (max-width: 600px) {
  .section__title{
    font-size: 1.3em;
  }

  .apartado__title {
    font-size: 1rem;
  }
}
</style>