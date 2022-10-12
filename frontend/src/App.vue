<script setup>
import { RouterView } from 'vue-router'

import Header from './components/Header.vue';
</script>

<template>
  <Header></Header>
  <RouterView v-slot="{ Component, route}">
    <template v-if="Component">
      <Transition mode="out-in" :name="route.meta.transitionName">
        <KeepAlive :max="1" :include="/^search/ig">
          <component :is="Component"></component>
        </KeepAlive>
      </Transition>
    </template>
  </RouterView>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  background-color: black;
}

:root {
  --color-icons: hsl(0, 0%, 17%); /* Very Dark Gray*/
  --color-secondary: hsl(0, 0%, 59%); /* Dark Gray*/
}

/*Transitions beetwen components*/
.slide-left-enter-active,
.slide-left-leave-active {
  position: absolute;
  transition: all 0.6s ease;
}

.slide-left-enter-from {
  left: 100%;
}

.slide-left-enter-to {
  left: 0%;
}

.slide-left-leave-from {
  transform: scale(1);
}

.slide-left-leave-to {
  transform: scale(0.8);
}

.slide-right-enter-active,
.slide-right-leave-active {
  position: absolute;
  transition: all 0.6s ease;
}
.slide-right-enter-from {
  left: -100%;
}
.slide-right-enter-to {
  left: 0%;
}
.slide-right-leave-from {
  transform: scale(1);
}
.slide-right-leave-to {
  transform: scale(0.8);
}
</style>