<template>
  <header class="navbar navbar-expand-md navbar-dark navbar-overlap">
    <div class="container-xl">
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbar-menu">
        <span class="navbar-toggler-icon" />
      </button>
      <a href="." class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pr-0 pr-md-3">
        <img :src="logoImage" alt="giregi.rip" class="navbar-brand-image">
      </a>
      <div class="navbar-nav flex-row order-md-last">
        <div v-if="useNotification" class="nav-item dropdown d-none d-md-flex mr-3">
          <a href="#" class="nav-link px-0" data-toggle="dropdown" tabindex="-1">
            <i class="far fa-bell" />
            <span v-if="notification" class="badge bg-red" />
          </a>
          <div class="dropdown-menu dropdown-menu-right dropdown-menu-card">
            <div class="card">
              <div class="card-body">
                {{ notification }}
              </div>
            </div>
          </div>
        </div>
        <div class="nav-item dropdown">
          <a v-if="!isElectron" href="#" class="nav-link d-flex lh-1 text-reset p-0" data-toggle="dropdown">
            <span class="avatar" :style="`background-image: url(${profileImage})`" />
            <div class="d-none d-xl-block pl-2">
              <div>{{ username }}</div>
            </div>
          </a>
          <div class="dropdown-menu dropdown-menu-right">
            <h6 class="dropdown-header d-xl-none">{{ username }}</h6>
            <a class="dropdown-item" href="#" @click="logout">
              <i class="fas fa-sign-out-alt" />Logout
            </a>
          </div>
        </div>
      </div>
      <div id="navbar-menu" class="collapse navbar-collapse">
        <div class="d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center">
          <ul class="navbar-nav">
            <li class="nav-item">
              <router-link to="/" :class="`nav-link ${isActive('/')}`">
                <i class="fas fa-home" />
                <span class="nav-link-title">Home</span>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/archives/my" :class="`nav-link ${isActive('/archives/my')}`">
                <i class="fas fa-archive" />
                <span class="nav-link-title">내 아카이브 목록</span>
              </router-link>
            </li>
            <!-- li class="nav-item dropdown active">
              <a class="nav-link dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
                <i class="fas fa-archive" />
                <span class="nav-link-title">아카이브 목록</span>
              </a>
              <ul class="dropdown-menu dropdown-menu-columns  dropdown-menu-columns-2">
                <li>
                  <a class="dropdown-item" href="./empty.html">
                    Empty page
                  </a>
                </li>
              </ul>
            </li -->
          </ul>
          <div v-if="showSearch" class="ml-md-auto pl-md-4 py-2 py-md-0 mr-md-4 order-first order-md-last flex-grow-1 flex-md-grow-0">
            <form action="." method="get">
              <div class="input-icon">
                <span class="input-icon-addon">
                  <i class="fas fa-search" />
                </span>
                <input type="text" class="form-control form-control-dark" placeholder="Search…">
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import md5 from 'md5';
import { mapState } from 'vuex';
import ApiClient, { API } from 'api/client';
import logoImage from 'assets/logo-white-horizontal.png';

export default {
  name: 'SiteHeader',
  mixins: [ApiClient],
  props: {
    notification: {
      type: String,
      required: false,
      default: '',
    },
    useNotification: {
      type: Boolean,
      require: false,
      default: false,
    },
    useSearch: {
      type: Boolean,
      require: false,
      default: false,
    },
  },
  data: () => ({ logoImage }),
  computed: {
    ...mapState({
      username: (state) => {
        return state.user.email;
      },
    }),
    isElectron() {
      return __ELECTRON__;
    },
    profileImage() {
      return `http://www.gravatar.com/avatar/${md5(this.username)}`;
    },
    showSearch() {
      return false;
    },
  },
  methods: {
    isActive(path) {
      return this.$route.path === path ? 'active' : '';
    },
    async logout() {
      const UserApi = this.getApi(API.USER);
      await UserApi.logout();
      await this.$store.dispatch('user/set', {});
    },
  },
};
</script>

<style scoped lang="scss">
/* fix - dropdown indicator rotating */
.dropdown-toggle::after {
  display: inline-block;
  margin-left: .255em;
  vertical-align: .255em;
  content: "";
  border-top: .3em solid;
  border-right: .3em solid transparent;
  border-bottom: 0;
  border-left: .3em solid transparent;
  transform: none;
}
i {
  &.fas {
    margin-right: 5px;
  }
  &.far {
    margin-right: 5px;
  }
}
.navbar-brand-image {
  height: 28px;
}
</style>
