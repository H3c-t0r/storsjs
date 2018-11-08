import { shallowMount, mount } from '@vue/test-utils';
import HeaderedInput from '@/components/HeaderedInput.vue';

describe('HeaderedInput.vue', () => {
	
	it('renders correctly with default props', () => {

    	const wrapper = shallowMount(HeaderedInput);

		expect(wrapper).toMatchSnapshot();
	});

	it('renders correctly with isMultiline props', () => {

    	const wrapper = shallowMount(HeaderedInput, {
			propsData: { isMultiline: true }
		});

		expect(wrapper).toMatchSnapshot();
		expect(wrapper.contains("textarea")).toBe(true);
		expect(wrapper.contains("input")).toBe(false);
	});

	it('renders correctly with size props', () => {
		let label = "testLabel";
		let width = "30px";
		let height = "20px";

    	const wrapper = shallowMount(HeaderedInput, {
			propsData: { label, width, height }
		});

		expect(wrapper.find("input").element.style.width).toMatch(width);
		expect(wrapper.find("input").element.style.height).toMatch(height);
		expect(wrapper.find(".labelContainer").text()).toMatch(label);
	});

	it('renders correctly with isOptional props', () => {

    	const wrapper = shallowMount(HeaderedInput, {
			propsData: {
				isOptional: true	
			}
		});

		expect(wrapper.find("h4").text()).toMatch("Optional");
	});
	  
	it('renders correctly with input error', () => {
		let error = "testError";

    	const wrapper = shallowMount(HeaderedInput, {
			propsData: {
				error	
			}
		});

		expect(wrapper).toMatchSnapshot();
		expect(wrapper.find(".labelContainer").text()).toMatch(error);
	});
  
  	it('emit setData on input correctly', () => {
		let testData = "testData";

		const wrapper = mount(HeaderedInput);

		wrapper.find("input").trigger('input');
		
		expect(wrapper.emitted("setData").length).toEqual(1);

		wrapper.vm.$emit('setData', testData);

		expect(wrapper.emitted("setData")[1][0]).toEqual(testData);
	});

});