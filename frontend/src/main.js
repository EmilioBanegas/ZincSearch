import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from 'pinia'
import axios from "axios";

axios.defaults.baseURL = "http://localhost:3000/api";


import { Quasar } from 'quasar'

// Import icon libraries
import '@quasar/extras/roboto-font/roboto-font.css'
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/fontawesome-v6/fontawesome-v6.css'

// Import Quasar css
import 'quasar/src/css/index.sass'

const store = createPinia()

const app = createApp(App);

app.use(router);
app.use(store);

app.use(Quasar, {
  plugins: {}, // import Quasar plugins and add here
});

app.mount("#app");
