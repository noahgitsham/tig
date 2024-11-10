{
  description = "Tig flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: let 
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
    common-deps = with pkgs; [ libcap go gcc ];
  in {

    devShell.x86_64-linux = pkgs.mkShell {
      buildInputs = with pkgs; [  ] ++ common-deps;
    };

    packages.x86_64-linux.tig = pkgs.buildGoPackage {
      goDeps = with pkgs; [  ] ++ common-deps;
    };
  };
}
