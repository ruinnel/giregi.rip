const state = {
  id: 1,
  email: 'email@domain.com',
  isAdmin: false,
};

// getters
const getters = {
  id: (state) => state.id,
  email: (state) => state.email,
  isAdmin: (state) => state.isAdmin,
};

// actions
const actions = {
  set ({ commit }, user) {
    const { id, email, isAdmin } = user;
    commit('id', id);
    commit('email', email);
    commit('isAdmin', isAdmin);
  },
};

// mutations
const mutations = {
  id (state, id) {
    state.id = id;
  },
  email (state, email) {
    state.email = email;
  },
  isAdmin (state, isAdmin) {
    state.isAdmin = isAdmin;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
