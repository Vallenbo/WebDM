import {get, post, postRaw} from "@/util/request"

const BASE_URL = "/api"

export function CreateContainer(conf) {
    return post(`${BASE_URL}/CreateContainer`, conf)
}

export function RunContainer(id) {
    return post(`${BASE_URL}/RunContainer?id=${id}`)
}

export function StopContainer(containerID) {
    return post(`${BASE_URL}/StopContainer?id=${containerID}`)
}

export function StopremoveContainer(containerID) {
    return post(`${BASE_URL}/StopremoveContainer?id=${containerID}`)
}

export function ListContainers() {
    return post(`${BASE_URL}/ListContainers`)
}

export function ContainerInspect(containerID) {
    return post(`${BASE_URL}/ContainerInspect?id=${containerID}`)
}

export function MonitorContainer() {
    return post(`${BASE_URL}/MonitorContainer`)
}

export function LogsContainer(containerID) {
    return post(`${BASE_URL}/LogsContainer?id=${containerID}`)
}

export function GetDockerVersion() {
    return post(`${BASE_URL}/GetDockerVersion`)
}

export function VolumesContainer(containerID) {
    return post(`${BASE_URL}/VolumesContainer?id=${containerID}`)
}

export function CreateImage(imageInput) {
    return post(`${BASE_URL}/CreateImage`, imageInput)
}

export function GetAllImages() {
    return post(`${BASE_URL}/GetAllImages`)
}

export function DeleteImage(imageidID) {
    return post(`${BASE_URL}/DeleteImage?id=${imageidID}`)
}

export function SearchImage(imageidName) {
    return post(`${BASE_URL}/SearchImage?imageName=${imageidName}`)
}

export function pullImage(imageidID) {
    return postRaw(`${BASE_URL}/PullImage?id=${imageidID}`)
}

export function dockerfilereturn() {
    return post(`${BASE_URL}/dockerfilereturn`)
}

export function monitor() {
    return post(`${BASE_URL}/monitor`)
}

export function containerInfo(id) {
    return post(`${BASE_URL}/ContainerInspect?id=${id}`)
}