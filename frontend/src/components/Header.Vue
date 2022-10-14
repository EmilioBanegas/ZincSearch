<template>
  <header id="header" class="header">
    <h1 class="header__title">Mamumo - Buscador</h1>
  </header>
</template>

<style scoped>
.header {
  position: relative;
  height: 20vh;
  max-width: 100vw;
  background-repeat: no-repeat;
  background-size: cover;
  background-image: url("./../assets/pattern-bg.png");
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  align-items: center;
  padding: 2vh 0;
  padding-bottom: 4vh;
}

.header__title {
  color: white;
  font-weight:700;
  max-height: 10vh;
  font-size:4em;
}

@media (max-width: 600px) {
  .header__title{
    font-size: 1.8em;
  }
}
</style>