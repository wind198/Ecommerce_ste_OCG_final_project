<template>
  <header  >
    <div class="announcement">
      <a href="#"
        >Up to 70% off RRP for The Style Sale! Use code STYLE for additional 10%
        off the collection</a
      >
    </div>
    <div class="main-info" @mouseleave="act_hideLogPanel" >
      <div class="customer-care">
        <a href="#" role="button" title="Customer Care">
          customer care<i class="far fa-comment"></i>
        </a>
      </div>
      <div class="logo"><a href="#">interior online</a></div>
      <div class="user-toolbar" >
        <a href="#" role="button" title="Utilities" class="util">
          <i class="fas fa-mug-hot"></i>
        </a>
        <a href="#" role="button" title="Search"
          ><i class="fas fa-search"></i
        ></a>
        <a href="#" role="button" title="Favourite" class="favourite fas"
          ><i class="far fa-heart"></i>
        </a>
        <a
          class="user"
          @mouseover="act_showLogPanel"
          href="#"
          role="button"
          title="User profile"
          ><i class="far fa-user-circle"></i></a
        ><a href="#" role="button" title="Shopping cart"
          ><i class="fas fa-shopping-cart"></i
        ></a>
        <div class="login-logout" v-show="displayLogPanel">
          <a v-if="!userlogged" class="login" href="#">Log in</a>
          <a v-if="userlogged" class="logout" href="#">Log out</a>
        </div>
      </div>
    </div>

    <nav>
      <ul>
        <li><a href="#">furniture</a></li>
        <li><a href="#">decor</a></li>
        <li><a href="#">lighting</a></li>
        <li><a href="#">outdoor & graden</a></li>
        <li><a href="#">decor styles</a></li>
        <li><a href="#">inspiration</a></li>
      </ul>
      <div class="pull-down">
        <div class="pull-bar"></div>
        <div class="handle">
          <i class="fas fa-long-arrow-alt-down"></i>
        </div>
      </div>
    </nav>

    <div class="slogan">
      <span>Shop our carefully curated selection of interior design décor</span>
    </div>
  </header>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  name: "Header",
  components: {},
  methods: {
    ...mapActions(["act_hideLogPanel","act_showLogPanel"]),
  },
  computed: {
    ...mapGetters(["userlogged", "displayLogPanel"]),
  },
};
</script>

<style lang="scss" scoped>
@media only screen and (max-width: 768px) {
  header {
    position: relative;
    .announcement {
      padding: 0.5rem 0;
      a {
        display: block;
        text-align: center;
        font-size: 0.7rem;
      }
    }
    .main-info {
      position: relative;
      padding: 0.75rem 0.5rem;
      display: grid;
      grid-template-columns: 33.3% 33.3% 33.3%;
      .customer-care {
        a {
          font-size: 0.9rem;
          i {
            margin-left: 0.3rem;
          }
        }
      }
      .user-toolbar {
        text-align: right;
        position: absolute;
        z-index: 20;
        right: 10px;
        top: 50%;
        transform: translateY(-18px);
        div.login-logout {
          display: block;
          text-align: center;
          a {
            text-align: center;
            background: white;
            color: black;
            text-transform: capitalize;
            margin: 0 0.1rem;
            border: 1px solid gray;
            border-radius: 5px 5px 0 0;
            line-height: 30px;
            font-size: 0.6rem;
            padding: 0.2rem 1rem;
          }
        }
        > a {
          display: none;
          font-size: 1.1rem;
        }
        > a.util {
          display: block;
        }
        &:hover {
          > a {
            display: inline-block;
            margin-right: 0.5rem;
          }
          > a.util {
            display: none;
          }
        }
      }
    }
    nav {
      position: absolute;
      width: 100%;
      &:hover {
        ul {
          display: block;
        }
        .pull-down {
          display: none;
        }
      }
      ul {
        display: none;
        border-bottom: 1px solid $bg-color1;

        li {
          padding: 0;
          font-size: 0.9rem;
          a {
            padding: 0.75rem 0;
            transition: all 0.1s linear;
          }
          &:hover {
            a {
              background: white;
              color: black;
            }
          }
        }
      }
      .pull-down {
        position: relative;
        background: none;
        .pull-bar {
          height: 2px;
          border-top: 4px double black;
        }
        .handle {
          text-align: center;
          i {
            top: 3px;
            position: absolute;
            animation: handle-animation 3s linear 1s infinite;
          }
          @keyframes handle-animation {
            0% {
              transform: rotateZ(0deg);
            }
            30% {
              transform: rotateZ(360deg);
              top: 3px;
              opacity: 1;
            }

            80% {
              top: 40px;
              transform: rotateZ(360deg);
              opacity: 0;
            }
            99% {
              top: 40px;
              transform: rotateZ(360deg);
              opacity: 0;
            }
          }
        }
      }
    }
    .slogan {
      margin-top: 1.5rem;
      font-size: 0.8rem;
    }
  }
}

header {
  margin-bottom: 2rem;
  .announcement {
    background-color: black;
    a {
      &:hover {
        text-decoration: underline;
      }
    }
  }
  .main-info {
    > div {
      line-height: 33px;
    }
    background-color: $bg-color1;
    color: white;
    border-bottom: 1px solid white;

    .customer-care {
      a {
        text-transform: capitalize;
      }
    }
    .logo {
      text-align: center;
      a {
        display: block;
        text-transform: uppercase;
        &::first-letter {
          font-size: 150%;
        }
      }
    }
  }
  nav {
    z-index: 10;
    ul {
      background-color: $bg-color1;
      li {
        display: block;
        text-align: center;
        text-transform: uppercase;
        a {
          display: block;
        }
      }
    }
  }
  .slogan {
    width: 100%;
    text-align: center;
    span {
      display: inline-block;
      margin: 0 auto;
      border-bottom: 1px solid black;
    }
  }
}
</style>