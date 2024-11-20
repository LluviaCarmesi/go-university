import type IRole from "../../interfaces/IRole";
import getUserByToken from "./getUserByToken";
import getCookie from "../../utilities/getCookie";

export default async function checkCurrentUser() {
    const token = getCookie("token");

    const currentRole: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    }
    if (!token) {
        window.open("/login", "_self");
        return currentRole;
    }
    const userByTokenResponse = await getUserByToken(token);
    if (userByTokenResponse.isSuccessful) {
        switch (userByTokenResponse.user.Role) {
            case "student":
                currentRole.isStudent = true;
                break;
            case "professor":
                currentRole.isProfessor = true;
                break;
            case "admin":
                currentRole.isAdmin = true;
                break;
        }
    }
    return currentRole;
}