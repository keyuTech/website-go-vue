import request from "@/api/index"

export const getSettings = () => request.get("/settings")
