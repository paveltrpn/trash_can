
#ifndef __geometry_h_
#define __geometry_h_


class geometry_c {
	public:
		float	*data;
		int		vcount;

		geometry_c(): data{nullptr}, vcount{0} {};
		~geometry_c();

		virtual void load_from_file(char *fname);
};

class static_geometry_c : geometry_c {
	public:
		float	*data;
		int		vcount;

		static_geometry_c() : data{nullptr}, vcount{0} {};
		~static_geometry_c();

		void load_from_file(char *fname);
};

class semi_static_geometry_c : geometry_c {
	public:
		float	*data;
		float	*collision_data;

		int		vcount;
		int		collision_vcount;
		
		semi_static_geometry_c(): data{nullptr}, vcount{0} {};
		~semi_static_geometry_c();

		void load_from_file(char *fname);
};

#endif