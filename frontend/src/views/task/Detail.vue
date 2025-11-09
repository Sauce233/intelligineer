<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { useSessionStore } from '@/stores/session';
import { request } from '@/utils/axios';
import type { Task } from '@/utils/tables';
import { formatDate } from '@/utils/utils';
import { Back, Edit } from '@element-plus/icons-vue';
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import {useRouteParams} from '@vueuse/router';

const id = useRouteParams('id')
const task = ref<Task | null>(null)
const route = useRoute()
const session = useSessionStore()
request<Task>('GET', `/tasks/${route.params.id}`).then(res => task.value = res)

</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div v-if="task" class="card text-lg flex items-center flex-wrap gap-2">
      <el-button :icon="Back" @click="$router.push('/tasks')">{{$t('back')}}</el-button>
      <span>{{$t('task')}}: {{task.name}}</span>
      <span class="link" @click="$router.push(`/projects/${task.project.id}`)">
        {{task.project.name}}
      </span>
      <tag :color="task.status.color">{{task.status.name}}</tag>
      <el-button :icon="Edit" class="ms-auto" @click="$router.push(`/tasks/${id}/edit`)">
        {{$t('taskDetail.editButton')}}
      </el-button>
    </div>
    <div v-if="task" class="flex gap-4 flex-wrap">
      <div v-for="v, k in {
            startTime: formatDate(task.startTime),
            endTime: formatDate(task.endTime),
            createdAt: formatDate(task.createdAt),
            updatedAt: formatDate(task.updatedAt) }" :key="k"
        class="flex gap-2 items-center"
      >
        <div class="text-subtle">{{$t(k)}}</div>
        <div>{{v}}</div>
      </div>
    </div>
    <div v-if="task" class="bg-subtle p-4 rounded-2xl columns-1 md:columns-2" v-html="session.md.render(task.profile)">
    </div>

    <div class="card flex gap-4">
      <router-link
        v-for="r in [
          [`/tasks/${id}`, $t('taskMaterial')],
          [`/tasks/${id}/records`, $t('record')],
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
