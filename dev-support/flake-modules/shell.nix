{ ... }:
{
  perSystem = { pkgs, ... }: {
    devShells.default =
      with pkgs;
      let
        clang-tools = clang-tools_16;
        llvmPackages = llvmPackages_16;
        # 1.21.6 is not in upstream
        go = go_1_21.overrideAttrs (_: rec {
          version = "1.21.6";
          src = fetchurl {
            url = "https://go.dev/dl/go${version}.src.tar.gz";
            hash = "sha256-Ekkmpi5F942qu67bnAEdl2MxhqM8I4/8HiUyDAIEYkg=";
          };
        });
        nodePackages = nodePackages_latest;
      in
      mkShellNoCC {
        packages =
          [
            direnv
            nix-direnv

            gitAndTools.git-extras
            gitAndTools.pre-commit
            nodePackages."@commitlint/cli"
            tokei

            treefmt
            nixpkgs-fmt
            deadnix
            shfmt
            shellcheck
            nodePackages.prettier
            nodePackages.sql-formatter
            taplo
            hclfmt
            clang-tools
            codespell

            docker
            docker-buildx

            gnugrep
            gawk
            nettools

            cmake
            pkg-config
            llvmPackages.clang

            go
            gotools
            golangci-lint
            govulncheck

            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
          ];

        shellHook = ''
          export PATH=$PWD/dev-support/bin:$PATH
          export GOBIN=$PWD/bin;
          git config --global url."git@github.com:".insteadOf "https://github.com"
        '';

        COMMITLINT_PRESET = "${nodePackages."@commitlint/config-conventional"}/lib/node_modules/@commitlint/config-conventional";
        GOPRIVATE = "github.com/chelpis/*";
      };
  };
}
