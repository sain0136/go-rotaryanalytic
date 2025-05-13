import { ref } from "vue";
import { authApi } from "../lib/api";

const isLoggedIn = ref(false);
const loggingout = ref(false);

export function useAuth() {
  if (sessionStorage.getItem("isLoggedIn") === "true") {
    isLoggedIn.value = true;
  }

  if (sessionStorage.getItem("loggingout") === "true") {
    loggingout.value = true;
  }

  const login = () => {
    isLoggedIn.value = true;
    sessionStorage.setItem("isLoggedIn", "true");
  };

  /**
   * Sets the loggingout state. This is used to let the whole app know that
   * we are currently logging out. This is useful for showing loading
   * indicators in the UI.
   *
   * @param {boolean} isLoggingOut - Whether we are currently logging out
   */
  const setLoggingOut = (isLoggingOut) => {
    loggingout.value = isLoggingOut;
    sessionStorage.setItem("loggingout", JSON.stringify(isLoggingOut));
  };

  const logout = async () => {
    try {
      await authApi.logout();
      loggingout.value = true;
      setLoggingOut(true);
      sessionStorage.removeItem("isLoggedIn");
      window.location.href = "/login";
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };

  const checkAuth = () => {
    isLoggedIn.value = sessionStorage.getItem("isLoggedIn") === "true";
    return isLoggedIn.value;
  };

  return {
    isLoggedIn,
    loggingout,
    login,
    logout,
    checkAuth,
    setLoggingOut,
  };
}
