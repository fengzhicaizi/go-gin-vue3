import { defineStore } from 'pinia';
import { getAuthByTokenApi } from '@/apis/system/auth';
import type { AuthStateType } from './auth.d';

export const useAuthStore = defineStore({
	id: 'auth',
	state: (): AuthStateType => ({
		auth: null,
		role: null,
		menu: null,
	}),

	getters: {
		authGetter: state => state.auth,
		roleGetter: state => state.role,
		menusGetter: state => state.menu,
	},

	actions: {
		async getAuth() {
			try {
				const res = await getAuthByTokenApi();
				const { data } = res;
				if (data) {
					const r = data.data;
					this.auth = r.auth;
					this.role = r.role;
					this.menu = r.menu;
				}
			} catch (err) {
				console.log(err);
			}
		},
	},
});
