

export function getBackendBaseURL(): URL {
    if (import.meta.env.MANGAWEB_BACKEND_URL !== undefined && import.meta.env.MANGAWEB_BACKEND_URL != "") {
        return new URL(import.meta.env.MANGAWEB_BACKEND_URL)
    }

    return new URL("http://localhost:8972");
}