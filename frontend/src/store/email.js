import { defineStore } from 'pinia'

export const useEmailStore = defineStore('email-store', {
  state: () => {
    return {
      email: null,
    }
  },

  getters: {
    getEmail(state) {
      return state.email;
    }
  },

  actions: {
    setEmail(newEmail) {
      this.email = newEmail;
    }    
  }
})