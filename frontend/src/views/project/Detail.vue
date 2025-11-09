<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {useSessionStore} from '@/stores/session';
import {request} from '@/utils/axios';
import type {Project} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {Back, Delete, Edit} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useRouteParams} from '@vueuse/router';

const id = useRouteParams('id')
const project = ref<Project | null>(null)
const session = useSessionStore()
request<Project>('GET', `/projects/${id.value}`).then(res => project.value = res)

</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div v-if="project" class="card text-lg flex items-center flex-wrap gap-2">
      <el-button :icon="Back" @click="$router.push('/projects')">{{$t('back')}}</el-button>
      <span>{{$t('project')}}: {{project.name}}</span>
      <span class="link" @click="$router.push(`/clients/${project.client.id}`)">{{project.client.name}}</span>
      <tag :color="project.status.color">{{project.status.name}}</tag>
      <el-button :icon="Edit" class="ms-auto" circle @click="$router.push(`/projects/${id}/edit`)"></el-button>
      <el-button :icon="Delete" circle @click="request('DELETE', `/projects/${id}`).then(v => v && $router.back())">
      </el-button>
    </div>
    <div v-if="project" class="flex gap-4 flex-wrap">
      <div v-for="v, k in {
           offer: project.offer,
           budget: project.budget,
           createdAt: formatDate(project.createdAt),
           updatedAt: formatDate(project.updatedAt) }" :key="k"
        class="flex gap-2 items-center"
      >
        <div class="text-subtle">{{$t(k)}}</div>
        <div>{{v}}</div>
      </div>
    </div>
    <div v-if="project" class="bg-subtle p-4 rounded-2xl columns-1 md:columns-2" v-html="session.md.render(project.profile)">
    </div>

    <div class="card flex gap-4">
      <router-link
        v-for="r in [
          [`/projects/${id}`, $t('attachment')],
          [`/projects/${id}/tasks`, $t('task')],
          [`/projects/${id}/invoices`, $t('invoice')],
        ]"
        :key="r[0]"
        :to="r[0]"
        exact-active-class="font-bold"
      >
        {{r[1]}}
      </router-link>
    </div>

    <router-view>
    </router-view>
  </div>

</template>
