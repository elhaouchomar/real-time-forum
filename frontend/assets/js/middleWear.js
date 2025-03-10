import { apiRequest } from "./apiRequest.js"
export async function MiddleWear(){
    const response = await apiRequest("checker")
    return response
}
