<?php get_header(); ?>
  <div class="wrap">
    <!-- ///page-title -->
    <div class="page-head">
      <h1 class="page-head__title">文庫中央鍼灸整骨院ブログ</h1>
    </div>
    <!-- mainvisual/// -->
    <?php if(have_posts()) :
      while(have_posts()) :
      the_post();
      remove_filter('the_content', 'wpautop');
    ?>
    <!-- ///content -->
    <div class="content blog-wrap">
    <?php if ( wp_is_mobile() ) : ?>
      <?php get_sidebar(); ?>
    <?php endif; ?>
      <div class="mr40">
        <article>
          <div class="blog-article">
            <div class="blog-headline">
              <p class="blog-headline__date">
              <?php echo get_the_time('Y'); ?>.<?php echo get_the_time('n'); ?>.<?php echo get_the_time('j'); ?>. <?php echo get_the_time('l'); ?> <?php echo get_the_time('G:i'); ?>
              </p>
              <h2 class="blog-headline__title"><?php the_title(); ?></h2>
            </div>
            <div class="blog-content"><?php the_content(); ?></div>
          </div>
        </article>
      <?php endwhile; else: ?>

      <?php endif; ?>
      <div class="blog-pager">
        <?php if (get_next_post()):?>
        <div class="next"><?php next_post_link('%link','<i class="fas fa-angle-double-left mr5"></i>next</a>',FALSE,'5'); ?></div>
        <?php endif; ?>
        <?php if (get_previous_post()):?>
        <div class="back"><?php previous_post_link('%link','<i class="fas fa-angle-double-right mr5"></i>back</a>',FALSE,'5'); ?></div>
        <?php endif; ?>
      </div>
    </div>
    <?php if ( !wp_is_mobile() ) : ?>
      <?php get_sidebar(); ?>
    <?php endif; ?>
  </div>
<?php get_footer();