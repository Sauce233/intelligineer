<script setup lang="ts">
import { useSessionStore } from '@/stores/session';
import { request } from '@/utils/axios';
import type { Client } from '@/utils/tables';
import { formatDate } from '@/utils/utils';
import { Back, Delete, Edit } from '@element-plus/icons-vue';
import { reactive, ref } from 'vue';
import { useRouteParams } from '@vueuse/router';
import ClientEditor, {type Form} from '@/components/editor/ClientEditor.vue';

const id = useRouteParams('id')
const client = ref<Client | null>(null)
const session = useSessionStore()
const load = async () => client.value = await request('GET', `/clients/${id.value}`)
load()
const isDialogOpen = ref(false)
const clientForm = reactive<Form>({
  name: '', address: '', profile: '',
})
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div v-if="client" class="card text-lg flex items-center flex-wrap gap-2">
      <el-button :icon="Back" @click="$router.back">{{$t('back')}}</el-button>
      <span class="text-xl">{{$t('clientDetail.title')}}: {{client.name}}</span>
      <el-button :icon="Edit" class="ms-auto" circle @click="isDialogOpen = true; Object.assign(clientForm, client)">
      </el-button>
      <el-button :icon="Delete" circle
        @click="request('DELETE', `/clients/${id}`).then(v => v && $router.push('/clients'))"
      ></el-button>
    </div>
    <div class="text-xl">{{$t('clientDetail.infoTitle')}}</div>
    <div v-if="client" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="v in [
          ['name', client.name, ''],
          ['updatedAt', formatDate(client.updatedAt), ''],
          ['createdAt', formatDate(client.createdAt), ''],
          ['address', client.address, 'col-span-2'],
        ]" :key="v[0]"
        :class="['flex gap-2 items-center', v[2]]"
      >
        <div class="text-subtle">{{$t(v[0])}}</div>
        <div>{{v[1]}}</div>
      </div>
    </div>

    <div class="text-xl">{{$t('clientDetail.profileTitle')}}</div>

    <div
      v-if="client"
      class="bg-subtle p-4 rounded-2xl columns-1 md:columns-2"
      v-html="session.md.render(client.profile)"
    />

    <div class="card flex gap-4">
      <router-link
        v-for="r in [
          [`/clients/${id}`, $t('contact')],
          [`/clients/${id}/projects`, $t('project')],
        ]"
        :key="r[0]"
        :to="r[0]"
        exact-active-class="font-bold"
      >
        {{r[1]}}
      </router-link>
    </div>

    <router-view />

  </div>

  <el-dialog v-model="isDialogOpen" class="w-dialog" :title="$t('clientDetail.clientDialogTitle')">
    <client-editor v-model="clientForm">
    </client-editor>
    <template #footer>
      <el-button type="primary" @click="request('PATCH', `/clients/${id}`, clientForm).then(ok => ok && (load(), isDialogOpen = false))">
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
