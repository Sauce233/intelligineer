<script setup lang="ts">
import type {Form} from '@/components/editor/PurchaseEditor.vue';
import PurchaseEditor from '@/components/editor/PurchaseEditor.vue';
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Purchase} from '@/utils/tables';
import {Back, Edit} from '@element-plus/icons-vue';
import {useRouteQuery} from '@vueuse/router';
import {reactive, ref} from 'vue';
import {useI18n} from 'vue-i18n';
import {useRoute} from 'vue-router';

const id = useRouteQuery('id')
const route = useRoute()

const purchase = ref<Purchase | null>(null)
const load = () => request<Purchase>('GET', `/purchases/${route.params.id}`).then(res => purchase.value = res)
load()

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', profile: '', statusId: 0,
})
const { t } = useI18n({})
</script>

<template>
  <div v-if="purchase" class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="card flex items-center gap-2 flex-wrap">
      <el-button :icon="Back" @click="$router.back">{{$t('back')}}</el-button>
      <div>{{purchase.name}}</div>
      <div class="text-subtle link" @click="$router.push(`/suppliers/${purchase.supplier.id}`)">
        {{purchase.supplier.name}}
      </div>
      <tag :color="purchase.status.color">{{purchase.status.name}}</tag>
      <el-button :icon="Edit" class="ms-auto" @click="isDialogOpen = true">
        {{$t('edit')}}
      </el-button>
    </div>
    <div class="bg-subtle rounded-2xl p-4">{{purchase.profile}}</div>
    <div class="font-bold text-xl">{{$t('totalPrice')}}: {{purchase.materials.reduce((sum, v) => sum + v.price * v.quantity, 0)}}</div>
    <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="m in purchase.materials" :key="m.material.id" @click="$router.push(`/materials/${m.material.id}`)" class="card cursor-pointer">
        <div class="space-x-2">
          <span>{{m.material.name}}</span>
          <span class="text-subtle">{{m.material.category.name}}</span>
        </div>
        <div class="flex justify-between">
          <div>{{m.quantity}} {{m.material.unit}}</div>
          <div>¥ {{m.price}} / {{m.material.unit}}</div>
          <div>¥ {{m.quantity * m.price}}</div>
        </div>
      </div>
    </div>
  </div>

  <el-dialog v-model="isDialogOpen" :title="t('dialogTitle')">
    <purchase-editor v-model="form" />
    <template #footer>
      <el-button @click="request('PATCH', `/purchases/${id}`, form).then(ok => ok && load())">
        {{t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
