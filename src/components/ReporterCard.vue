<template>
  <vs-card>
    <vs-row slot="header" class="header">
      <div class="avatar"><vs-avatar slot="avatar" /></div>
      <div class="info">
        <div>
          <h3>
            {{ reporter.name }}
            <span v-if="hasMyMemo && !nowEdit" class="memo text-danger">{{ myMemo }}</span>
            <input v-if="nowEdit" v-model="memo" type="text" class="memo-input" />
            <vs-icon
              :icon="nowEdit ? 'fa-save' : 'fa-edit'"
              icon-pack="fas"
              class="memo-save"
              @click.native="onIconClick"
            />
          </h3>
        </div>
        <div class="meta">
          <vs-icon icon-pack="far" icon="fa-building" />
          {{ agencyName }}
          <vs-icon icon-pack="far" icon="fa-envelope" />
          {{ email }}
          <i class="far fa-clock" />
          {{ formatDateTime(reporter.createdAt) }}
        </div>
      </div>
    </vs-row>
    <div class="content">
      <slot />
      <vs-collapse v-if="hasExtraAgencies">
        <vs-collapse-item>
          <div slot="header">소속 이력</div>
          <ul>
            <li v-for="(agency, agencyIdx) in extraAgencies" :key="agencyIdx">
              {{ agency.name }}
            </li>
          </ul>
        </vs-collapse-item>
      </vs-collapse>
    </div>
    <div slot="footer">
      <vs-row class="buttons">
        <vs-button
          type="border"
          icon-pack="far"
          icon="fa-thumbs-up"
          color="primary"
          class="like"
          @click.stop="reaction(true)"
        >
          {{ formatNumber(reporter.like) }}
        </vs-button>
        <vs-button
          type="border"
          icon-pack="far"
          icon="fa-thumbs-up fa-rotate-180"
          color="danger"
          class="unlike"
          @click.stop="reaction(false)"
        >
          {{ formatNumber(reporter.unlike) }}
        </vs-button>
      </vs-row>
    </div>
  </vs-card>
</template>

<script>
import { last, size, slice, isEmpty, get, isString } from 'lodash';

export default {
  name: 'ReporterCard',
  props: {
    reporter: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      inputMemo: null,
      nowEdit: false,
    };
  },
  computed: {
    memo: {
      get() {
        return isString(this.inputMemo) ? this.inputMemo : get(this.reporter, 'myMemo.content');
      },
      set(val) {
        this.inputMemo = val;
      },
    },
    hasMyMemo() {
      return !isEmpty(this.reporter.myMemo);
    },
    myMemo() {
      return get(this.reporter, 'myMemo.content');
    },
    agencyName() {
      return (last(this.reporter.agencies) || {}).name;
    },
    email() {
      return (last(this.reporter.agencies) || {}).email || '알수없음';
    },
    hasExtraAgencies() {
      return size(this.reporter.agencies) > 1;
    },
    extraAgencies() {
      return slice(this.reporter.agencies, 0, size(this.reporter.agencies) - 1);
    },
  },
  methods: {
    onClick() {
      console.log('click');
    },
    reaction(isLike) {
      this.$emit('reaction', { isLike, reporterId: this.reporter.id });
    },
    async onIconClick() {
      if (this.nowEdit) {
        const id = get(this.reporter, 'myMemo.id');
        this.$emit('memo', { id, memo: this.inputMemo });
        this.nowEdit = false;
      } else {
        this.nowEdit = true;
      }
    },
  },
};
</script>

<style scoped lang="scss">
.header {
  > div.avatar {

  }
  > div {
    > i {
      font-size: 1em;
    }
  }
  > div.info {
    > div > h3 {
      > .memo {
        margin-left: 5px;
      }
      > .memo-input {
        width: unset;
        height: unset;
        display: inline;
        padding: 0 0.5em;
        margin-left: 5px;
      }
      > .memo-save {
        margin-left: 5px;
        cursor: pointer;
      }
      display: inline;
      font-weight: 400;
    }
    > div.meta {
      display: flex;
      align-items: center;
      > i {
        margin-right: 3px;
        margin-left: 5px;
      }
      > i:first-child {
        margin-right: 3px;
        margin-left: 0px;
      }
    }
    font-size: 0.9em;
    margin-left: 5px;
  }
}
.buttons {
  justify-content: flex-end;
  > button {
    width: 100px;
    height: unset;
    &.unlike {
      margin-left: 5px;
    }
  }
}
</style>
