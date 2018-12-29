import { usersModule } from '@/store/modules/users';
import { changePasswordRequest, deleteAccountRequest, getUserRequest, updateAccountRequest } from '@/api/users';
import { USER_MUTATIONS } from '@/store/mutationConstants';

const { SET_USER_INFO, REVERT_TO_DEFAULT_USER_INFO, UPDATE_USER_INFO } = usersModule.mutations;
const { updateAccount, changePassword, deleteAccount, getUser } = usersModule.actions;
const { user, userName } = usersModule.getters;

describe('mutations', () => {
	it('Set user info', () => {
		const state = {
			user: {
				firstName: '',
				lastName: '',
				email: ''
			}
		};

		const user = {
			firstName: 'firstName',
			lastName: 'lastName',
			email: 'email'
		};

		SET_USER_INFO(state, user);

		expect(state.user.email).toBe('email');
		expect(state.user.firstName).toBe('firstName');
		expect(state.user.lastName).toBe('lastName');
	});

	it('clear user info', () => {
		const state = {
			user: {
				firstName: 'firstName',
				lastName: 'lastName',
				email: 'email',
			}
		};

		REVERT_TO_DEFAULT_USER_INFO(state);

		expect(state.user.email).toBe('');
		expect(state.user.firstName).toBe('');
		expect(state.user.lastName).toBe('');
	});

	it('Update user info', () => {
		const state = {
			user: {
				firstName: '',
				lastName: '',
				email: ''
			}
		};

		const user = {
			firstName: 'firstName',
			lastName: 'lastName',
			email: 'email'
		};

		UPDATE_USER_INFO(state, user);

		expect(state.user.email).toBe('email');
		expect(state.user.firstName).toBe('firstName');
		expect(state.user.lastName).toBe('lastName');
	});
});

describe('actions', () => {
	it('success update account', async () => {
		updateAccountRequest = jest.fn().mockReturnValue({
			isSuccess: true, data: {
				firstName: 'firstName',
				lastName: 'lastName',
				email: 'email',
			}
		});
		const commit = jest.fn();
		const user = {
			firstName: '',
			lastName: '',
			email: ''
		};
		await updateAccount({commit}, user);

		expect(commit).toHaveBeenCalledWith(USER_MUTATIONS.UPDATE_USER_INFO, {
			firstName: 'firstName',
			lastName: 'lastName',
			email: 'email'
		});
	});

	it('error update account', async () => {
		updateAccountRequest = jest.fn().mockReturnValue({
			isSuccess: false
		});
		const commit = jest.fn();
		const user = {
			firstName: '',
			lastName: '',
			email: ''
		};
		await updateAccount({commit}, user);

		expect(commit).toHaveBeenCalledTimes(0);
	});

	it('password change', async () => {
		changePasswordRequest = jest.fn().mockReturnValue({isSuccess: true});
		const commit = jest.fn();
		const updatePasswordModel = {oldPassword: 'o', newPassword: 'n'};
		const requestResponse = await changePassword({commit}, updatePasswordModel);

		expect(requestResponse.isSuccess).toBeTruthy();
	});

	it('delete account', async () => {
		deleteAccountRequest = jest.fn().mockReturnValue({isSuccess: true});
		const commit = jest.fn();
		const password = '';

		const requestResponse = await deleteAccount(commit, password);
		expect(requestResponse.isSuccess).toBeTruthy();
	});

	it('success get user', async () => {
		getUserRequest = jest.fn().mockReturnValue({
			isSuccess: true,
			data: {
				firstName: '',
				lastName: '',
				email: ''
			}
		});

		const commit = jest.fn();
		const requestResponse = await getUser({commit});
		expect(requestResponse.isSuccess).toBeTruthy();
	});

	it('error get user', async () => {
		getUserRequest = jest.fn().mockReturnValue({isSuccess: false});

		const commit = jest.fn();
		const requestResponse = await getUser({commit});
		expect(requestResponse.isSuccess).toBeFalsy();
	});
});

describe('getters', () => {
	it('user model', function () {
		const state = {
			user: {
				firstName: 'firstName',
				lastName: 'lastName',
				email: 'email',
			}
		};

		const retrievedUser = user(state);

		expect(retrievedUser.firstName).toBe('firstName');
		expect(retrievedUser.lastName).toBe('lastName');
		expect(retrievedUser.email).toBe('email');
	});
	it('user name', function () {
		const state = {
			user: {
				firstName: 'John',
				lastName: 'Doe'
			}
		};

		const retrievedUserName = userName(state);

		expect(retrievedUserName).toBe('John Doe');
	});
});
