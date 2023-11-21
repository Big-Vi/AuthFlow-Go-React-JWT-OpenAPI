const SERVER_URL = `http://localhost:8000/api`;

export const ApiEndpoints = {
    //Management
    LOGIN: `${SERVER_URL}/user/login`,
    SIGNUP: `${SERVER_URL}/user/signup`,
    AUTHSTATUS: `${SERVER_URL}/user/auth-status`,
    LOGOUT: `${SERVER_URL}/user/logout`
}
