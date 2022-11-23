import { ElNotification } from 'element-plus'

export const notifyError = (e: any) => {
    ElNotification({
        title: e.response.data,
        type: 'error'
    })
}

export const notifySuccess = (msg: string) => {
    ElNotification({
        title: msg,
        type: 'success'
    })
}