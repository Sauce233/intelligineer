<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { useSessionStore } from '@/stores/session';
import { request } from '@/utils/axios';
import type { Task } from '@/utils/tables';
import { formatDate } from '@/utils/utils';
import { Back, Edit } from '@element-plus/icons-vue';
import { reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import {useRouteParams} from '@vueuse/router';
import {useI18n} from 'vue-i18n';
import TaskEditor, {type Form} from '@/components/editor/TaskEditor.vue';

const id = useRouteParams('id')
const task = ref<Task | null>(null)
const route = useRoute()
const isDialogOpen = ref(false)
const session = useSessionStore()
const load = () => request<Task>('GET', `/tasks/${route.params.id}`).then(res => task.value = res)
load()

const { t } = useI18n({})
const form = reactive<Form>({
  name: '', profile: '', time: [new Date(), new Date()], statusId: 0,
})

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
      <el-button :icon="Edit" class="ms-auto" @click="isDialogOpen = true">
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

  <el-dialog v-model="isDialogOpen" :title="t('dialogTitle')">
    <task-editor v-model="form" />
    <template #footer>
      <el-button @click="request('PATCH', `/tasks/${id}`, form).then(ok => ok && load())">
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
